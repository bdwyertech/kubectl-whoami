// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package autofill

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAutofill(in *jlexer.Lexer, out *TriggerParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "fieldId":
			(out.FieldID).UnmarshalEasyJSON(in)
		case "frameId":
			(out.FrameID).UnmarshalEasyJSON(in)
		case "card":
			if in.IsNull() {
				in.Skip()
				out.Card = nil
			} else {
				if out.Card == nil {
					out.Card = new(CreditCard)
				}
				(*out.Card).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAutofill(out *jwriter.Writer, in TriggerParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"fieldId\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.FieldID))
	}
	if in.FrameID != "" {
		const prefix string = ",\"frameId\":"
		out.RawString(prefix)
		out.String(string(in.FrameID))
	}
	{
		const prefix string = ",\"card\":"
		out.RawString(prefix)
		if in.Card == nil {
			out.RawString("null")
		} else {
			(*in.Card).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TriggerParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAutofill(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TriggerParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAutofill(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TriggerParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAutofill(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TriggerParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAutofill(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAutofill1(in *jlexer.Lexer, out *CreditCard) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "number":
			out.Number = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "expiryMonth":
			out.ExpiryMonth = string(in.String())
		case "expiryYear":
			out.ExpiryYear = string(in.String())
		case "cvc":
			out.Cvc = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAutofill1(out *jwriter.Writer, in CreditCard) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"number\":"
		out.RawString(prefix[1:])
		out.String(string(in.Number))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"expiryMonth\":"
		out.RawString(prefix)
		out.String(string(in.ExpiryMonth))
	}
	{
		const prefix string = ",\"expiryYear\":"
		out.RawString(prefix)
		out.String(string(in.ExpiryYear))
	}
	{
		const prefix string = ",\"cvc\":"
		out.RawString(prefix)
		out.String(string(in.Cvc))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreditCard) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAutofill1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreditCard) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAutofill1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreditCard) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAutofill1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreditCard) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAutofill1(l, v)
}
