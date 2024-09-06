package core

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/nichirinto/nichirin/framework/lib/datetime"
	"net/http"
)

func (r *Router[T, I]) Auth() *Router[T, I] {
	r.AuthGuard = true

	return r
}

func (r *Router[TI, TO]) Register(app *Nichirin, httpMethod HttpMethod, url string, fn func(*Context[TI, TO]) *Res[TO]) *Router[TI, TO] {
	r.Url = url
	r.Method = httpMethod
	r.Name = r.Url
	r.Handler = fn

	return r.Attach(app)
}

func handle[T any, I any](w http.ResponseWriter, req *http.Request, r *Router[T, I]) {
	w.Header().Set("Content-Type", "application/json")

	input := new(T)

	if r.Method == Get {
		query := req.URL.Query()

		queryMap := make(map[string]interface{})
		for key, values := range query {
			if len(values) > 0 {
				queryMap[key] = values[0]
			}
		}

		queryData, err := json.Marshal(queryMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(queryData, input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		if err := json.NewDecoder(req.Body).Decode(input); err != nil {
			ThrowBadRequest(w, err)
			return
		}
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(input); err != nil {
		ThrowBadRequest(w, err)
		return
	}

	c := NewContext[T, I](req)
	c.Input = input
	c.Exception = nil
	c.Output = new(I)

	if r.Handler == nil {
		return
	}

	res := r.Handler(c)
	c.Timeline.Finished = datetime.Unix()

	if res.Err != nil {
		c.Logger.Error(res.Err.Err.Error())
	}

	SendResponse(w, res)
}

func (r *Router[T, I]) Attach(app *Nichirin) *Router[T, I] {
	if r.Method == Get {
		app.Engine.Get(r.Url, func(w http.ResponseWriter, req *http.Request) {
			handle(w, req, r)
		})

		return r
	}

	if r.Method == Put {
		app.Engine.Put(r.Url, func(w http.ResponseWriter, req *http.Request) {
			handle(w, req, r)
		})

		return r
	}

	if r.Method == Patch {
		app.Engine.Patch(r.Url, func(w http.ResponseWriter, req *http.Request) {
			handle(w, req, r)
		})

		return r
	}

	if r.Method == Delete {
		app.Engine.Delete(r.Url, func(w http.ResponseWriter, req *http.Request) {
			handle(w, req, r)
		})

		return r
	}

	app.Engine.Post(r.Url, func(w http.ResponseWriter, req *http.Request) {
		handle(w, req, r)
	})

	return r
}
