/*
 * Copyright 2023 Pius Alfred <me.pius1102@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software
 * and associated documentation files (the “Software”), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial
 * portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT
 * LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
 * WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"context"
	"encoding/json"
	werrors "github.com/piusalfred/whatsapp/errors"
	"github.com/piusalfred/whatsapp/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/piusalfred/whatsapp/webhooks"
)

type Middleware func(next http.Handler) http.Handler

// Wrap wraps a http.Handler with a middlewares
func Wrap(handler http.Handler, middlewares ...Middleware) http.Handler {
	// wraps backwards
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// StartTimeMiddleware is a middleware that adds the start time to the request context
func StartTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("StartTimeMiddleware")
		ctx := r.Context()
		ctx = context.WithValue(ctx, "startTime", time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// EndTimeMiddleware is a middleware that adds the end time to the request context
func EndTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("EndTimeMiddleware")
		ctx := r.Context()
		ctx = context.WithValue(ctx, "endTime", time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

var _ webhooks.Hooks = (*handler)(nil)

type handler struct {
}

func (h *handler) OnMessageStatusChange(ctx context.Context, nctx *webhooks.NotificationContext, status *webhooks.Status) error {
	//TODO implement me
	panic("implement me")
}

func (h *handler) OnNotificationError(ctx context.Context, nctx *webhooks.NotificationContext, errors *werrors.Error) error {
	//TODO implement me
	panic("implement me")
}

func (h *handler) OnMessageReceived(ctx context.Context, nctx *webhooks.NotificationContext, message *webhooks.Message, hooks webhooks.MessageHooks) error {
	//TODO implement me
	panic("implement me")
}

var _ webhooks.MessageHooks = (*messageHooks)(nil)

type messageHooks struct{}

func (m messageHooks) OnMessageErrors(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, errors []*werrors.Error) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnTextMessageReceived(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, text *webhooks.Text) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnReferralMessageReceived(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, text *webhooks.Text, referral *webhooks.Referral) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnCustomerIDChange(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, customerID *webhooks.Identity) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnSystemMessage(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, system *webhooks.System) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnImageReceived(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, image *models.MediaInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnAudioReceived(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, audio *models.MediaInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnVideoReceived(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, video *models.MediaInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnDocumentReceived(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, document *models.MediaInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnStickerReceived(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, sticker *models.MediaInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnOrderReceived(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, order *webhooks.Order) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnButtonMessage(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, button *webhooks.Button) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnLocationReceived(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, location *models.Location) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnContactsReceived(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, contacts models.Contacts) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnMessageReaction(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, reaction *models.Reaction) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnUnknownMessageReceived(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, errors []*werrors.Error) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnProductEnquiry(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, text *webhooks.Text) error {
	//TODO implement me
	panic("implement me")
}

func (m messageHooks) OnInteractiveMessage(ctx context.Context, nctx *webhooks.NotificationContext, mctx *webhooks.MessageContext, interactive *webhooks.Interactive) error {
	//TODO implement me
	panic("implement me")
}

func (h *handler) HandleError(ctx context.Context, writer http.ResponseWriter, request *http.Request, err error) error {
	if err != nil {
		os.Stdout.WriteString(err.Error())
	}

	os.Stdout.WriteString("error is nil")
	return nil
}

func (h *handler) HandleEvent(ctx context.Context, writer http.ResponseWriter, request *http.Request, notification *webhooks.Notification) error {
	os.Stdout.WriteString("HandleEvent")
	jsonb, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonb)
	return nil
}

/*
	curl -X POST --location "http://localhost:8080/webhooks" \
		   -H "Content-Type: application/json" \
		   -d "{
		         "object": "whatsapp_business_account",
		         "entry": [{
		           "id": "WHATSAPP_BUSINESS_ACCOUNT_ID",
		           "changes": [{
		             "value": {
		               "messaging_product": "whatsapp",
		               "metadata": {
		                 "display_phone_number": "PHONE_NUMBER",
		                 "phone_number_id": "PHONE_NUMBER_ID"
		               },
		               "contacts": [{
		                 "profile": {
		                  	"name": "NAME"
		                 },
		                 "wa_id": "PHONE_NUMBER_ID"
		               }],
		               "messages": [{
		                 "from": "PHONE_NUMBER",
		                 "id": "wamid.ID",
		                 "timestamp": "TIMESTAMP",
		                 "text": {
		                   "body": "MESSAGE_BODY"
		                 },
		                 "type": "text"
		               }]
		             },
		             "field": "messages"
		           }]
		         }]
		       }"
*/
func main() {
	// Create a new handler
	//handler := &handler{}
	ls := webhooks.NewEventListener()
	middlewares := []Middleware{
		StartTimeMiddleware,
		EndTimeMiddleware,
	}
	finalHandler := Wrap(ls.Handle(), middlewares...)
	mux := http.NewServeMux()
	mux.Handle("/webhooks", finalHandler)

	// Create a new server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start the server
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
