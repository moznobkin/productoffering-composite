/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/antihax/optional"

	client "github.com/moznobkin/productoffering-composite/life-client"
)

const regexp = `"((\d+).+\sбесплатно,\s)?(.*\s)(\d+)\s(рубля|рублей|рубль)\s(в\sдень|месяц|год|неделю)"`

func GetInstalledBase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetQualifiedCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetQualifiedProductOfferings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var filterMap map[string]CategoryRef

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	msisdn := ""
	keys, ok := r.URL.Query()["msisdn"]
	if ok && len(keys) > 0 {
		msisdn = keys[0]
	}
	if len(data) > 1 {
		reader := bytes.NewReader(data)

		var filter []CategoryRef
		err := json.NewDecoder(reader).Decode(&filter)
		if err != nil {
			log.Println("Failed to load filter categories")
		} else {
			if len(filter) > 0 {
				filterMap = make(map[string]CategoryRef, len(filter))
				for _, cat := range filter {
					filterMap[cat.Name] = cat
				}
			}
		}

	}
	cats, resp, err := getCategoriesQ(msisdn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		fmt.Println(resp)
		return
	}
	// cats, err := getCategories(msisdn)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	result, err := mapCategories2Qualification(cats, filterMap)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
}

func parsePay(pay string) (string, float32, error) {
	return "Day", 0, nil

}

func mapLabels(labels []client.CategoryOffersLabels) []OfferLabels {
	result := make([]OfferLabels, len(labels))
	for i, l := range labels {
		result[i] = OfferLabels{
			Name:        l.Labelname,
			Description: l.Labelcolor,
		}
	}
	return result
}

func mapProductOfferingPrice(subs []client.Subscription) []ProductOfferingPrice {
	result := make([]ProductOfferingPrice, len(subs))

	for i, s := range subs {
		RecurringChargePeriodType, price, err := parsePay(s.Pay)
		if err != nil {
			panic(err)
		}
		result[i] = ProductOfferingPrice{
			Description:               s.Description,
			IsBundle:                  false,
			LifecycleStatus:           "Active",
			Name:                      s.Heading,
			PriceType:                 "Recurring Charge",
			RecurringChargePeriodType: RecurringChargePeriodType,
			PricingLogicAlgorithm:     []PricingLogicAlgorithm{PricingLogicAlgorithm{Name: s.Pay, Description: s.Pay}},
			Price:                     &Money{Unit: "RUB", Value: price},
		}
	}
	return result
}

func mapCategories2Qualification(catOffers *[]client.CategoryOffers, filter map[string]CategoryRef) (*ProductOfferingQualificationResponse, error) {

	var containers []ProductOfferingQualificationResponseProductOfferingQualificationItemContainer = make([]ProductOfferingQualificationResponseProductOfferingQualificationItemContainer, 0)

	for _, cat := range *catOffers {
		var ok bool
		if filter != nil {
			_, ok = filter[cat.Categoryname]

		}
		if ok || filter == nil {
			for _, offer := range cat.Offers {
				p := ProductOfferingQualificationResponseProductOfferingQualificationItemContainer{
					QualificationItemResult: "Qualified",
					ProductOfferingQualificationItem: &ProductOfferingQualificationItem{
						BaseType: "object",

						Type_: "ProductOfferingQualificationItem",

						Action: "Qualify",
						Id:     strconv.FormatInt(offer.Id, 10),
						Category: &CategoryRef{
							Name: cat.Categoryname,
						},

						ProductOffering: &ProductOffering{
							Attachment: []Attachment{
								Attachment{
									Name: "miniature",
									Url:  offer.Miniature,
								},
								Attachment{
									Name: "header",
									Url:  offer.Header,
								},
							},
							Category: []CategoryRef{
								CategoryRef{
									Name: cat.Categoryname,
								},
							},
							Channel: []ChannelRef{
								ChannelRef{
									Name: "MLK",
								},
							},
							Description: offer.Shortdescription,
							IsBundle:    false,

							IsSellable: true,

							LifecycleStatus: "Active",

							Name: offer.Name,

							OfferLabels: mapLabels(offer.Labels),

							ProductOfferingPrice: mapProductOfferingPrice(offer.Subscriptions),

							SpecCharValueUse: []ConfigurableSpecificationCharacteristicValueUse{
								ConfigurableSpecificationCharacteristicValueUse{
									SpecificationType: "ProductSpecification",
									Description:       "Приоритет",
									MaxCardinality:    1,
									MinCardinality:    1,
									Name:              "priority",
									SpecCharacteristicValue: []SpecificationCharacteristicValue{
										SpecificationCharacteristicValue{
											IsDefault:     true,
											UnitOfMeasure: "NA",
											Value:         strconv.FormatInt(int64(offer.Priority), 10),
											ValueType:     "Integer",
										},
									},
								},
							},
							StatusReason: "Qualification",
							Version:      "1.0",
						},
					},
				}
				containers = append(containers, p)
			}
		}
	}
	p := ProductOfferingQualificationResponse{
		ProductOfferingQualificationItemContainer: containers,
	}
	return &p, nil
}

func getCategories(msisdn string) (*[]client.CategoryOffers, error) {
	if len(msisdn) != 10 && len(msisdn) != 0 {
		return nil, errors.New("Bad request")
	}
	fs, err := os.Open("./examples/json/offerslist.json")
	if err != nil {
		panic(err)
	}
	var p []client.CategoryOffers
	err = json.NewDecoder(fs).Decode(&p)

	if err != nil {
		return nil, err
	}
	if len(msisdn) == 0 {
		for _, cat := range p {
			for i, _ := range cat.Offers {
				cat.Offers[i].Subscriptions = nil
			}
		}
	}
	return &p, nil
}
func getCategoriesQ(msisdn string) (*[]client.CategoryOffers, *http.Response, error) {
	if len(msisdn) != 10 && len(msisdn) != 0 {
		return nil, nil, errors.New("Bad request")
	}
	cfg := &client.Configuration{
		BasePath:      "https://mf-offers-core.quantum-a.io", //"http://localhost:8081/life",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	cl := client.NewAPIClient(cfg)
	optMsidn := optional.NewString("7" + msisdn)
	ctx := context.WithValue(context.Background(), client.ContextBasicAuth, client.BasicAuth{UserName: "admin", Password: "Fe1muePh"})
	offers, resp, err := cl.ServiceAPIApi.GetOffers(ctx, &client.ServiceAPIApiGetOffersOpts{Msisdn: optMsidn})
	if err != nil {
		return nil, resp, err
	}

	return &offers.Category, resp, nil
}
