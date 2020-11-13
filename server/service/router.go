package service

import (
	dtype "RiYu/data/type"
	"RiYu/server/grammar/adjective"
	"RiYu/server/grammar/tense"
	"RiYu/server/grammar/verb"
	"RiYu/server/log"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var routers = []*router{
	{
		Method: http.MethodGet,
		Path:   "/ok",
		Handler: func(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
			defer func() {
				if err := recover(); err != nil {
					log.PanicError(err)
				}
			}()
			if _, err := writer.Write([]byte("可")); err != nil {
				panic(err)
			}
		},
	},
	{
		Method: http.MethodGet,
		Path:   "/getone",
		Handler: func(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
			defer func() {
				if err := recover(); err != nil {
					log.PanicError(err)
				}
			}()
			var (
				result   = getone()
				jsonByte []byte
				err      error
			)
			if jsonByte, err = json.Marshal(result); err != nil {
				panic(err)
			}
			if _, err = writer.Write(jsonByte); err != nil {
				panic(err)
			}
		},
	},
	{
		// 需要请求以 Content-Type:multipart/form-data 提交
		Method: http.MethodPost,
		Path:   "/verbforms",
		Handler: func(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
			defer func() {
				if err := recover(); err != nil {
					log.PanicError(err)
				}
			}()
			var (
				varbForm *verb.Forms
				jsonByte []byte
				err      error
			)
			if err = req.ParseMultipartForm(0); err != nil {
				panic(err)
			}
			if varbForm, err = verb.GetVerbForms(req.PostFormValue("词"), req.PostFormValue("标")); err != nil {
				panic(err)
			}
			if jsonByte, err = json.Marshal(varbForm); err != nil {
				panic(err)
			}
			if _, err = writer.Write(jsonByte); err != nil {
				panic(err)
			}
		},
	},
	{
		// 需要请求以 Content-Type:multipart/form-data 提交
		Method: http.MethodPost,
		Path:   "/adjforms",
		Handler: func(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
			defer func() {
				if err := recover(); err != nil {
					log.PanicError(err)
				}
			}()
			var (
				adjForm  *adjective.Forms
				jsonByte []byte
				err      error
			)
			if err = req.ParseMultipartForm(0); err != nil {
				panic(err)
			}
			if adjForm, err = adjective.GetAdjForms(req.PostFormValue("词"), req.PostFormValue("标")); err != nil {
				panic(err)
			}
			if jsonByte, err = json.Marshal(adjForm); err != nil {
				panic(err)
			}
			if _, err = writer.Write(jsonByte); err != nil {
				panic(err)
			}
		},
	},
	{
		Method: http.MethodPost,
		Path:   "/tense",
		Handler: func(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
			defer func() {
				if err := recover(); err != nil {
					log.PanicError(err)
				}
			}()
			var (
				tenseForm *tense.Forms
				jsonByte  []byte
				err       error
			)
			if err = req.ParseMultipartForm(0); err != nil {
				panic(err)
			}
			if tenseForm, err = tense.GetTenseForms(req.PostFormValue("词"), req.PostFormValue("标")); err != nil {
				panic(err)
			}
			if jsonByte, err = json.Marshal(tenseForm); err != nil {
				panic(err)
			}
			if _, err = writer.Write(jsonByte); err != nil {
				panic(err)
			}
		},
	},
	{
		Method: http.MethodGet,
		Path:   "/getone/tense",
		Handler: func(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
			defer func() {
				if err := recover(); err != nil {
					log.PanicError(err)
				}
			}()
			var (
				forms    *tense.Forms
				jsonByte []byte
				err      error
			)
			if forms, err = getoneTense(); err != nil {
				panic(err)
			}
			if jsonByte, err = json.Marshal(forms); err != nil {
				panic(err)
			}
			if _, err = writer.Write(jsonByte); err != nil {
				panic(err)
			}
		},
	},
	{
		Method: http.MethodPost,
		Path:   "/search",
		Handler: func(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
			defer func() {
				if err := recover(); err != nil {
					log.PanicError(err)
				}
			}()
			var (
				words    []*dtype.Words
				jsonByte []byte
				err      error
			)
			if err = req.ParseMultipartForm(0); err != nil {
				panic(err)
			}
			if words, err = search(req.PostForm); err != nil {
				panic(err)
			}
			if jsonByte, err = json.Marshal(words); err != nil {
				panic(err)
			}
			if _, err = writer.Write(jsonByte); err != nil {
				panic(err)
			}
		},
	},
}
