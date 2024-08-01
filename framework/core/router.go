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

func (r *Router[T, I]) Path(url string) *Router[T, I] {
	r.Url = url
	r.Name = r.Url

	return r
}

func (r *Router[TI, TO]) Handle(fn func(*Context[TI]) *Res[TO]) *Router[TI, TO] {
	r.Handler = fn

	return r
}

func (r *Router[T, I]) Attach(app *Nichirin) *Router[T, I] {
	app.Engine.Post(r.Url, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		input := new(T)

		if err := json.NewDecoder(req.Body).Decode(input); err != nil {
			ThrowBadRequest(w, err)
			return
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		if err := validate.Struct(req.Body); err != nil {
			ThrowBadRequest(w, err)
			return
		}

		c := NewContext[T](req)
		c.Input = input

		if r.Handler == nil {
			return
		}

		res := r.Handler(c)
		c.Timeline.Finished = datetime.Unix()

		if res.Err != nil {
			c.Logger.Error(res.Err.Err.Error())
		}

		SendResponse(w, res)
	})

	return r
}
