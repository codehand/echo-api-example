package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"

	"github.com/codehand/echo-restful-crud-api-example/handler"
	"github.com/codehand/echo-restful-crud-api-example/middlewares"
	"github.com/codehand/echo-restful-crud-api-example/types"
)

// TestGetAllProducts is func test get all product - all case
func TestGetAllProducts(t *testing.T) {
	e := echo.New()
	e.Validator = middlewares.InitCustomValidator()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	tests := []struct {
		name    string
		data    []*types.Product
		isError types.PayloadStatus
	}{
		{
			name:    "Case 1: get all products",
			data:    nil,
			isError: types.OkStatus,
		},
	}

	for _, test := range tests {
		assert.NotEmpty(t, test.name, "Name testing invalid")
		req := httptest.NewRequest(echo.GET, "/", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetPath("/api/v1/products")
		err := handler.GetProducts(c)
		if test.isError.HasError() {
			assert.Error(t, err, test.name)
			assert.NotEqual(t, http.StatusOK, res.Code, test.name)
			var es *types.PayloadStatus
			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err, "Parse body error")
			assert.NoError(t, json.Unmarshal(body, &es), "Parse body error")
			assert.NotEmpty(t, es.Code, "Code error empty")
			assert.NotEmpty(t, es.Message, "Message empty")
		} else {
			assert.NoError(t, err, test.name)
			assert.Equal(t, http.StatusOK, res.Code, test.name)
			var data []*types.Product
			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err, "Parse body error")
			assert.NoError(t, json.Unmarshal(body, &data), "Parse body error")
			assert.NotEqual(t, 0, len(data), test.name)
			for _, item := range data {
				assert.NotEmpty(t, item.Name, test.name)
				assert.NotEmpty(t, item.ImageClosed, test.name)
				assert.NotEmpty(t, item.ImageOpen, test.name)
				assert.NotEmpty(t, item.Description, test.name)
				assert.NotEmpty(t, item.Story, test.name)
				assert.NotEmpty(t, item.AllergyInfo, test.name)
				assert.NotEmpty(t, item.DietaryCertifications, test.name)
				assert.NotEqual(t, 0, item.ProductID, test.name)
			}
		}
	}
}

// TestGetOneProduct is func test get one product - all case
func TestGetOneProduct(t *testing.T) {
	e := echo.New()
	e.Validator = middlewares.InitCustomValidator()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	tests := []struct {
		name    string
		input   int
		output  *types.Product
		isError types.PayloadStatus
	}{
		{
			name:  "Case 1: get one products - OK",
			input: 1,
			output: &types.Product{
				Name:                  "Vanilla Toffee Bar Crunch",
				Description:           "Vanilla Ice Cream with Fudge-Covered Toffee Pieces",
				SourcingValues:        []string{"Non-GMO", "Cage-Free Eggs", "Fairtrade", "Responsibly Sourced Packaging", "Caring Dairy"},
				AllergyInfo:           "may contain wheat, peanuts and other tree nuts",
				DietaryCertifications: "Kosher",
				ProductID:             "1",
			},
			isError: types.OkStatus,
		},
		{
			name:    "Case 2: get one products - FAIL",
			input:   0,
			output:  nil,
			isError: types.ParseStatus("NOT_FOUND", "not found"),
		},
		{
			name:    "Case 3: get one products - FAIL",
			input:   -1,
			output:  nil,
			isError: types.ParseStatus("NOT_FOUND", "not found"),
		},
		{
			name:    "Case 4: get one products - FAIL",
			input:   999999999,
			output:  nil,
			isError: types.ParseStatus("NOT_FOUND", "not found"),
		},
	}

	for _, test := range tests {
		assert.NotEmpty(t, test.name, "Name testing invalid")
		req := httptest.NewRequest(echo.GET, "/", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetPath("/api/v1/products/:id")
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprintf("%d", test.input))
		err := handler.GetProduct(c)
		if test.isError.HasError() {
			assert.NoError(t, err, test.name)
			assert.NotEqual(t, http.StatusOK, res.Code, test.name)
			var es *types.PayloadStatus
			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err, "Parse body error")
			assert.NoError(t, json.Unmarshal(body, &es), "Parse body error")
			assert.NotEmpty(t, es.Code, "Code error empty")
			assert.NotEmpty(t, es.Message, "Message empty")
			assert.Equal(t, test.isError.Code, es.Code, "Code not match")
		} else {
			assert.NoError(t, err, test.name)
			assert.Equal(t, http.StatusOK, res.Code, test.name)
			var data *types.Product
			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err, "Parse body error")
			assert.NoError(t, json.Unmarshal(body, &data), "Parse body error")
			assert.NotEmpty(t, data.Name, test.name)
			assert.NotEmpty(t, data.ImageClosed, test.name)
			assert.NotEmpty(t, data.ImageOpen, test.name)
			assert.NotEmpty(t, data.Description, test.name)
			assert.NotEmpty(t, data.Story, test.name)
			assert.NotEmpty(t, data.AllergyInfo, test.name)
			assert.NotEmpty(t, data.DietaryCertifications, test.name)
			assert.NotEqual(t, 0, data.ProductID, test.name)
			assert.Equal(t, test.output.Name, data.Name, test.name)
			assert.Equal(t, test.output.Description, data.Description, test.name)
			assert.Equal(t, test.output.ProductID, data.ProductID, test.name)
			assert.Equal(t, test.output.AllergyInfo, data.AllergyInfo, test.name)
			assert.Equal(t, test.output.DietaryCertifications, data.DietaryCertifications, test.name)
			for _, item := range test.output.SourcingValues {
				assert.Contains(t, data.SourcingValues, item, test.name)
			}
		}
	}
}

// TestCreateOneProduct is func test create one product - all case
func TestCreateOneProduct(t *testing.T) {
	e := echo.New()
	e.Validator = middlewares.InitCustomValidator()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	tests := []struct {
		name    string
		input   []byte
		output  *types.Product
		isError types.PayloadStatus
	}{
		{
			name:  "Case 1: create one products - OK",
			input: []byte(`{"name":"Vanilla Toffee Bar Crunch","image_closed":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing.png","image_open":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing-open.png","description":"Vanilla Ice Cream with Fudge-Covered Toffee Pieces","story":"Vanilla What Bar Crunch? We gave this flavor a new name to go with the new toffee bars we’re using as part of our commitment to source Fairtrade Certified and non-GMO ingredients. We love it and know you will too!","sourcing_values":["Non-GMO","Cage-Free Eggs","Fairtrade","Responsibly Sourced Packaging","Caring Dairy"],"ingredients":["cream","skim milk","liquid sugar","water","sugar","coconut oil","egg yolks","butter","vanilla extract","almonds","cocoa (processed with alkali)","milk","soy lecithin","cocoa","natural flavor","salt","vegetable oil","guar gum","carrageenan"],"allergy_info":"may contain wheat, peanuts and other tree nuts","dietary_certifications":"Kosher"}`),
			output: &types.Product{
				Name:                  "Vanilla Toffee Bar Crunch",
				Description:           "Vanilla Ice Cream with Fudge-Covered Toffee Pieces",
				SourcingValues:        []string{"Non-GMO", "Cage-Free Eggs", "Fairtrade", "Responsibly Sourced Packaging", "Caring Dairy"},
				AllergyInfo:           "may contain wheat, peanuts and other tree nuts",
				DietaryCertifications: "Kosher",
			},
			isError: types.OkStatus,
		},
		{
			name:    "Case 2: create one products - missing field",
			input:   []byte(`{"name":"","image_closed":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing.png","image_open":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing-open.png","description":"Vanilla Ice Cream with Fudge-Covered Toffee Pieces","story":"Vanilla What Bar Crunch? We gave this flavor a new name to go with the new toffee bars we’re using as part of our commitment to source Fairtrade Certified and non-GMO ingredients. We love it and know you will too!","sourcing_values":["Non-GMO","Cage-Free Eggs","Fairtrade","Responsibly Sourced Packaging","Caring Dairy"],"ingredients":["cream","skim milk","liquid sugar","water","sugar","coconut oil","egg yolks","butter","vanilla extract","almonds","cocoa (processed with alkali)","milk","soy lecithin","cocoa","natural flavor","salt","vegetable oil","guar gum","carrageenan"],"allergy_info":"may contain wheat, peanuts and other tree nuts","dietary_certifications":"Kosher"}`),
			output:  nil,
			isError: types.ParseStatus("REQ_INVALID", "Vui lòng nhập giá trị Name"),
		},
		{
			name:    "Case 3: create one products - missing field",
			input:   []byte(`{"name":"name","image_closed":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing.png","image_open":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing-open.png","description":"Vanilla Ice Cream with Fudge-Covered Toffee Pieces","story":"Vanilla What Bar Crunch? We gave this flavor a new name to go with the new toffee bars we’re using as part of our commitment to source Fairtrade Certified and non-GMO ingredients. We love it and know you will too!","sourcing_values":["Non-GMO","Cage-Free Eggs","Fairtrade","Responsibly Sourced Packaging","Caring Dairy"],"ingredients":["cream","skim milk","liquid sugar","water","sugar","coconut oil","egg yolks","butter","vanilla extract","almonds","cocoa (processed with alkali)","milk","soy lecithin","cocoa","natural flavor","salt","vegetable oil","guar gum","carrageenan"],"allergy_info":"may contain wheat, peanuts and other tree nuts","dietary_certifications":""}`),
			output:  nil,
			isError: types.ParseStatus("REQ_INVALID", "Vui lòng nhập giá trị DietaryCertifications"),
		},
		{
			name:    "Case 4: create one products - incorrect body",
			input:   []byte(`{"name":11,"image_closed":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing.png","image_open":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing-open.png","description":"Vanilla Ice Cream with Fudge-Covered Toffee Pieces","story":"Vanilla What Bar Crunch? We gave this flavor a new name to go with the new toffee bars we’re using as part of our commitment to source Fairtrade Certified and non-GMO ingredients. We love it and know you will too!","sourcing_values":["Non-GMO","Cage-Free Eggs","Fairtrade","Responsibly Sourced Packaging","Caring Dairy"],"ingredients":["cream","skim milk","liquid sugar","water","sugar","coconut oil","egg yolks","butter","vanilla extract","almonds","cocoa (processed with alkali)","milk","soy lecithin","cocoa","natural flavor","salt","vegetable oil","guar gum","carrageenan"],"allergy_info":"may contain wheat, peanuts and other tree nuts","dietary_certifications":"test"}`),
			output:  nil,
			isError: types.ParseStatus("REQ_ERR", "Có lỗi xảy ra, vui lòng kiểm tra lại thông tin"),
		},
	}

	for _, test := range tests {
		assert.NotEmpty(t, test.name, "Name testing invalid")
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(test.input)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetPath("/api/v1/products")

		err := handler.CreateProduct(c)
		if test.isError.HasError() {
			assert.NoError(t, err, test.name)
			assert.NotEqual(t, http.StatusCreated, res.Code, test.name)
			var es *types.PayloadStatus
			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err, "Parse body error")
			assert.NoError(t, json.Unmarshal(body, &es), "Parse body error")
			assert.NotEmpty(t, es.Code, "Code error empty")
			assert.NotEmpty(t, es.Message, "Message empty")
			assert.Equal(t, test.isError.Code, es.Code, "Code not match")
			assert.Equal(t, test.isError.Message, es.Message, "Msg not match")
		} else {
			assert.NoError(t, err, test.name)
			assert.Equal(t, http.StatusCreated, res.Code, test.name)
			var data *types.Product
			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err, "Parse body error")
			assert.NoError(t, json.Unmarshal(body, &data), "Parse body error")
			assert.NotEmpty(t, data.Name, test.name)
			assert.NotEmpty(t, data.ImageClosed, test.name)
			assert.NotEmpty(t, data.ImageOpen, test.name)
			assert.NotEmpty(t, data.Description, test.name)
			assert.NotEmpty(t, data.Story, test.name)
			assert.NotEmpty(t, data.AllergyInfo, test.name)
			assert.NotEmpty(t, data.DietaryCertifications, test.name)
			assert.NotEqual(t, 0, data.ProductID, test.name)
			assert.Equal(t, test.output.Name, data.Name, test.name)
			assert.Equal(t, test.output.Description, data.Description, test.name)
			assert.Equal(t, test.output.AllergyInfo, data.AllergyInfo, test.name)
			assert.Equal(t, test.output.DietaryCertifications, data.DietaryCertifications, test.name)
			for _, item := range test.output.SourcingValues {
				assert.Contains(t, data.SourcingValues, item, test.name)
			}
		}
	}
}

// TestUpdateOneProduct is func test update one product - all case
func TestUpdateOneProduct(t *testing.T) {
	e := echo.New()
	e.Validator = middlewares.InitCustomValidator()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	tests := []struct {
		name    string
		id      int
		input   []byte
		output  *types.Product
		isError types.PayloadStatus
	}{
		{
			name:  "Case 1: update one products - OK",
			id:    4,
			input: []byte(`{"name":"Vanilla Toffee Bar Crunch","image_closed":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing.png","image_open":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing-open.png","description":"Vanilla Ice Cream with Fudge-Covered Toffee Pieces","story":"Vanilla What Bar Crunch? We gave this flavor a new name to go with the new toffee bars we’re using as part of our commitment to source Fairtrade Certified and non-GMO ingredients. We love it and know you will too!","sourcing_values":["Non-GMO","Cage-Free Eggs","Fairtrade","Responsibly Sourced Packaging","Caring Dairy"],"ingredients":["cream","skim milk","liquid sugar","water","sugar","coconut oil","egg yolks","butter","vanilla extract","almonds","cocoa (processed with alkali)","milk","soy lecithin","cocoa","natural flavor","salt","vegetable oil","guar gum","carrageenan"],"allergy_info":"may contain wheat, peanuts and other tree nuts","dietary_certifications":"Kosher"}`),
			output: &types.Product{
				Name:                  "Vanilla Toffee Bar Crunch",
				Description:           "Vanilla Ice Cream with Fudge-Covered Toffee Pieces",
				SourcingValues:        []string{"Non-GMO", "Cage-Free Eggs", "Fairtrade", "Responsibly Sourced Packaging", "Caring Dairy"},
				AllergyInfo:           "may contain wheat, peanuts and other tree nuts",
				DietaryCertifications: "Kosher",
			},
			isError: types.OkStatus,
		},
		{
			name:  "Case 2: update one products - OK",
			id:    4,
			input: []byte(`{"name":"name","image_closed":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing.png","image_open":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing-open.png","description":"Vanilla Ice Cream with Fudge-Covered Toffee Pieces","story":"Vanilla What Bar Crunch? We gave this flavor a new name to go with the new toffee bars we’re using as part of our commitment to source Fairtrade Certified and non-GMO ingredients. We love it and know you will too!","sourcing_values":["Non-GMO","Cage-Free Eggs","Fairtrade","Responsibly Sourced Packaging","Caring Dairy"],"ingredients":["cream","skim milk","liquid sugar","water","sugar","coconut oil","egg yolks","butter","vanilla extract","almonds","cocoa (processed with alkali)","milk","soy lecithin","cocoa","natural flavor","salt","vegetable oil","guar gum","carrageenan"],"allergy_info":"may contain wheat, peanuts and other tree nuts","dietary_certifications":"Kosher"}`),
			output: &types.Product{
				Name:                  "name",
				Description:           "Vanilla Ice Cream with Fudge-Covered Toffee Pieces",
				SourcingValues:        []string{"Non-GMO", "Cage-Free Eggs", "Fairtrade", "Responsibly Sourced Packaging", "Caring Dairy"},
				AllergyInfo:           "may contain wheat, peanuts and other tree nuts",
				DietaryCertifications: "Kosher",
			},
			isError: types.OkStatus,
		},
		{
			name:  "Case 3: update one products - OK",
			id:    4,
			input: []byte(`{"image_closed":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing.png","image_open":"/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing-open.png","description":"Vanilla Ice Cream with Fudge-Covered Toffee Pieces","story":"Vanilla What Bar Crunch? We gave this flavor a new name to go with the new toffee bars we’re using as part of our commitment to source Fairtrade Certified and non-GMO ingredients. We love it and know you will too!","sourcing_values":["Non-GMO","Cage-Free Eggs","Fairtrade","Responsibly Sourced Packaging","Caring Dairy"],"ingredients":["cream","skim milk","liquid sugar","water","sugar","coconut oil","egg yolks","butter","vanilla extract","almonds","cocoa (processed with alkali)","milk","soy lecithin","cocoa","natural flavor","salt","vegetable oil","guar gum","carrageenan"],"allergy_info":"may contain wheat, peanuts and other tree nuts","dietary_certifications":"Kosher"}`),
			output: &types.Product{
				Name:                  "name",
				Description:           "Vanilla Ice Cream with Fudge-Covered Toffee Pieces",
				SourcingValues:        []string{"Non-GMO", "Cage-Free Eggs", "Fairtrade", "Responsibly Sourced Packaging", "Caring Dairy"},
				AllergyInfo:           "may contain wheat, peanuts and other tree nuts",
				DietaryCertifications: "Kosher",
			},
			isError: types.OkStatus,
		},
	}

	for _, test := range tests {
		assert.NotEmpty(t, test.name, "Name testing invalid")
		req := httptest.NewRequest(echo.PUT, "/", strings.NewReader(string(test.input)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetPath("/api/v1/products/:id")
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprintf("%d", test.id))
		err := handler.UpdateProduct(c)
		if test.isError.HasError() {
			assert.NoError(t, err, test.name)
			assert.NotEqual(t, http.StatusOK, res.Code, test.name)
			var es *types.PayloadStatus
			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err, "Parse body error")
			assert.NoError(t, json.Unmarshal(body, &es), "Parse body error")
			assert.NotEmpty(t, es.Code, "Code error empty")
			assert.NotEmpty(t, es.Message, "Message empty")
			assert.Equal(t, test.isError.Code, es.Code, "Code not match")
			assert.Equal(t, test.isError.Message, es.Message, "Msg not match")
		} else {
			assert.NoError(t, err, test.name)
			assert.Equal(t, http.StatusOK, res.Code, test.name)
			var data *types.Product
			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err, "Parse body error")
			assert.NoError(t, json.Unmarshal(body, &data), "Parse body error")
			assert.NotEmpty(t, data.Name, test.name)
			assert.NotEmpty(t, data.ImageClosed, test.name)
			assert.NotEmpty(t, data.ImageOpen, test.name)
			assert.NotEmpty(t, data.Description, test.name)
			assert.NotEmpty(t, data.Story, test.name)
			assert.NotEmpty(t, data.AllergyInfo, test.name)
			assert.NotEmpty(t, data.DietaryCertifications, test.name)
			assert.NotEqual(t, 0, data.ProductID, test.name)
			assert.Equal(t, test.output.Name, data.Name, test.name)
			assert.Equal(t, test.output.Description, data.Description, test.name)
			assert.Equal(t, test.output.AllergyInfo, data.AllergyInfo, test.name)
			assert.Equal(t, test.output.DietaryCertifications, data.DietaryCertifications, test.name)
		}
	}
}
