// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: video.go
// DO NOT EDIT!

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Video) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Video) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if len(mj.Mimes) != 0 {
		buf.WriteString(`"mimes":`)
		if mj.Mimes != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Mimes {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Minduration != 0 {
		buf.WriteString(`"minduration":`)
		fflib.FormatBits2(buf, uint64(mj.Minduration), 10, mj.Minduration < 0)
		buf.WriteByte(',')
	}
	if mj.Maxduration != 0 {
		buf.WriteString(`"maxduration":`)
		fflib.FormatBits2(buf, uint64(mj.Maxduration), 10, mj.Maxduration < 0)
		buf.WriteByte(',')
	}
	if mj.Protocol != 0 {
		buf.WriteString(`"protocol":`)
		fflib.FormatBits2(buf, uint64(mj.Protocol), 10, mj.Protocol < 0)
		buf.WriteByte(',')
	}
	if len(mj.Protocols) != 0 {
		buf.WriteString(`"protocols":`)
		if mj.Protocols != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Protocols {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.W != 0 {
		buf.WriteString(`"w":`)
		fflib.FormatBits2(buf, uint64(mj.W), 10, mj.W < 0)
		buf.WriteByte(',')
	}
	if mj.H != 0 {
		buf.WriteString(`"h":`)
		fflib.FormatBits2(buf, uint64(mj.H), 10, mj.H < 0)
		buf.WriteByte(',')
	}
	if mj.Startdelay != 0 {
		buf.WriteString(`"startdelay":`)
		fflib.FormatBits2(buf, uint64(mj.Startdelay), 10, mj.Startdelay < 0)
		buf.WriteByte(',')
	}
	if mj.Linearity != 0 {
		buf.WriteString(`"linearity":`)
		fflib.FormatBits2(buf, uint64(mj.Linearity), 10, mj.Linearity < 0)
		buf.WriteByte(',')
	}
	if mj.Sequence != 0 {
		buf.WriteString(`"sequence":`)
		fflib.FormatBits2(buf, uint64(mj.Sequence), 10, mj.Sequence < 0)
		buf.WriteByte(',')
	}
	if len(mj.Battr) != 0 {
		buf.WriteString(`"battr":`)
		if mj.Battr != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Battr {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Maxextended != 0 {
		buf.WriteString(`"maxextended":`)
		fflib.FormatBits2(buf, uint64(mj.Maxextended), 10, mj.Maxextended < 0)
		buf.WriteByte(',')
	}
	if mj.Minbitrate != 0 {
		buf.WriteString(`"minbitrate":`)
		fflib.FormatBits2(buf, uint64(mj.Minbitrate), 10, mj.Minbitrate < 0)
		buf.WriteByte(',')
	}
	if mj.Maxbitrate != 0 {
		buf.WriteString(`"maxbitrate":`)
		fflib.FormatBits2(buf, uint64(mj.Maxbitrate), 10, mj.Maxbitrate < 0)
		buf.WriteByte(',')
	}
	if mj.Boxingallowed != 0 {
		buf.WriteString(`"boxingallowed":`)
		fflib.FormatBits2(buf, uint64(mj.Boxingallowed), 10, mj.Boxingallowed < 0)
		buf.WriteByte(',')
	}
	if len(mj.Playbackmethod) != 0 {
		buf.WriteString(`"playbackmethod":`)
		if mj.Playbackmethod != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Playbackmethod {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.Delivery) != 0 {
		buf.WriteString(`"delivery":`)
		if mj.Delivery != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Delivery {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Pos != 0 {
		buf.WriteString(`"pos":`)
		fflib.FormatBits2(buf, uint64(mj.Pos), 10, mj.Pos < 0)
		buf.WriteByte(',')
	}
	if len(mj.Companionad) != 0 {
		buf.WriteString(`"companionad":`)
		if mj.Companionad != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Companionad {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{
					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}
				}

			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.Api) != 0 {
		buf.WriteString(`"api":`)
		if mj.Api != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Api {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.Companiontype) != 0 {
		buf.WriteString(`"companiontype":`)
		if mj.Companiontype != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Companiontype {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.Ext) != 0 {
		buf.WriteString(`"ext":`)
		/* Falling back. type=openrtb.Extensions kind=map */
		err = buf.Encode(mj.Ext)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Videobase = iota
	ffj_t_Videono_such_key

	ffj_t_Video_Mimes

	ffj_t_Video_Minduration

	ffj_t_Video_Maxduration

	ffj_t_Video_Protocol

	ffj_t_Video_Protocols

	ffj_t_Video_W

	ffj_t_Video_H

	ffj_t_Video_Startdelay

	ffj_t_Video_Linearity

	ffj_t_Video_Sequence

	ffj_t_Video_Battr

	ffj_t_Video_Maxextended

	ffj_t_Video_Minbitrate

	ffj_t_Video_Maxbitrate

	ffj_t_Video_Boxingallowed

	ffj_t_Video_Playbackmethod

	ffj_t_Video_Delivery

	ffj_t_Video_Pos

	ffj_t_Video_Companionad

	ffj_t_Video_Api

	ffj_t_Video_Companiontype

	ffj_t_Video_Ext
)

var ffj_key_Video_Mimes = []byte("mimes")

var ffj_key_Video_Minduration = []byte("minduration")

var ffj_key_Video_Maxduration = []byte("maxduration")

var ffj_key_Video_Protocol = []byte("protocol")

var ffj_key_Video_Protocols = []byte("protocols")

var ffj_key_Video_W = []byte("w")

var ffj_key_Video_H = []byte("h")

var ffj_key_Video_Startdelay = []byte("startdelay")

var ffj_key_Video_Linearity = []byte("linearity")

var ffj_key_Video_Sequence = []byte("sequence")

var ffj_key_Video_Battr = []byte("battr")

var ffj_key_Video_Maxextended = []byte("maxextended")

var ffj_key_Video_Minbitrate = []byte("minbitrate")

var ffj_key_Video_Maxbitrate = []byte("maxbitrate")

var ffj_key_Video_Boxingallowed = []byte("boxingallowed")

var ffj_key_Video_Playbackmethod = []byte("playbackmethod")

var ffj_key_Video_Delivery = []byte("delivery")

var ffj_key_Video_Pos = []byte("pos")

var ffj_key_Video_Companionad = []byte("companionad")

var ffj_key_Video_Api = []byte("api")

var ffj_key_Video_Companiontype = []byte("companiontype")

var ffj_key_Video_Ext = []byte("ext")

func (uj *Video) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Video) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Videobase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Videono_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_Video_Api, kn) {
						currentKey = ffj_t_Video_Api
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffj_key_Video_Battr, kn) {
						currentKey = ffj_t_Video_Battr
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Video_Boxingallowed, kn) {
						currentKey = ffj_t_Video_Boxingallowed
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffj_key_Video_Companionad, kn) {
						currentKey = ffj_t_Video_Companionad
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Video_Companiontype, kn) {
						currentKey = ffj_t_Video_Companiontype
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffj_key_Video_Delivery, kn) {
						currentKey = ffj_t_Video_Delivery
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Video_Ext, kn) {
						currentKey = ffj_t_Video_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffj_key_Video_H, kn) {
						currentKey = ffj_t_Video_H
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffj_key_Video_Linearity, kn) {
						currentKey = ffj_t_Video_Linearity
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_Video_Mimes, kn) {
						currentKey = ffj_t_Video_Mimes
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Video_Minduration, kn) {
						currentKey = ffj_t_Video_Minduration
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Video_Maxduration, kn) {
						currentKey = ffj_t_Video_Maxduration
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Video_Maxextended, kn) {
						currentKey = ffj_t_Video_Maxextended
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Video_Minbitrate, kn) {
						currentKey = ffj_t_Video_Minbitrate
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Video_Maxbitrate, kn) {
						currentKey = ffj_t_Video_Maxbitrate
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_Video_Protocol, kn) {
						currentKey = ffj_t_Video_Protocol
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Video_Protocols, kn) {
						currentKey = ffj_t_Video_Protocols
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Video_Playbackmethod, kn) {
						currentKey = ffj_t_Video_Playbackmethod
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Video_Pos, kn) {
						currentKey = ffj_t_Video_Pos
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_Video_Startdelay, kn) {
						currentKey = ffj_t_Video_Startdelay
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Video_Sequence, kn) {
						currentKey = ffj_t_Video_Sequence
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffj_key_Video_W, kn) {
						currentKey = ffj_t_Video_W
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Ext, kn) {
					currentKey = ffj_t_Video_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Companiontype, kn) {
					currentKey = ffj_t_Video_Companiontype
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Api, kn) {
					currentKey = ffj_t_Video_Api
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Companionad, kn) {
					currentKey = ffj_t_Video_Companionad
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Video_Pos, kn) {
					currentKey = ffj_t_Video_Pos
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Delivery, kn) {
					currentKey = ffj_t_Video_Delivery
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Video_Playbackmethod, kn) {
					currentKey = ffj_t_Video_Playbackmethod
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Boxingallowed, kn) {
					currentKey = ffj_t_Video_Boxingallowed
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Maxbitrate, kn) {
					currentKey = ffj_t_Video_Maxbitrate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Minbitrate, kn) {
					currentKey = ffj_t_Video_Minbitrate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Maxextended, kn) {
					currentKey = ffj_t_Video_Maxextended
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Battr, kn) {
					currentKey = ffj_t_Video_Battr
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Video_Sequence, kn) {
					currentKey = ffj_t_Video_Sequence
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Linearity, kn) {
					currentKey = ffj_t_Video_Linearity
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Video_Startdelay, kn) {
					currentKey = ffj_t_Video_Startdelay
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_H, kn) {
					currentKey = ffj_t_Video_H
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_W, kn) {
					currentKey = ffj_t_Video_W
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Video_Protocols, kn) {
					currentKey = ffj_t_Video_Protocols
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Protocol, kn) {
					currentKey = ffj_t_Video_Protocol
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Maxduration, kn) {
					currentKey = ffj_t_Video_Maxduration
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Video_Minduration, kn) {
					currentKey = ffj_t_Video_Minduration
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Video_Mimes, kn) {
					currentKey = ffj_t_Video_Mimes
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Videono_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Video_Mimes:
					goto handle_Mimes

				case ffj_t_Video_Minduration:
					goto handle_Minduration

				case ffj_t_Video_Maxduration:
					goto handle_Maxduration

				case ffj_t_Video_Protocol:
					goto handle_Protocol

				case ffj_t_Video_Protocols:
					goto handle_Protocols

				case ffj_t_Video_W:
					goto handle_W

				case ffj_t_Video_H:
					goto handle_H

				case ffj_t_Video_Startdelay:
					goto handle_Startdelay

				case ffj_t_Video_Linearity:
					goto handle_Linearity

				case ffj_t_Video_Sequence:
					goto handle_Sequence

				case ffj_t_Video_Battr:
					goto handle_Battr

				case ffj_t_Video_Maxextended:
					goto handle_Maxextended

				case ffj_t_Video_Minbitrate:
					goto handle_Minbitrate

				case ffj_t_Video_Maxbitrate:
					goto handle_Maxbitrate

				case ffj_t_Video_Boxingallowed:
					goto handle_Boxingallowed

				case ffj_t_Video_Playbackmethod:
					goto handle_Playbackmethod

				case ffj_t_Video_Delivery:
					goto handle_Delivery

				case ffj_t_Video_Pos:
					goto handle_Pos

				case ffj_t_Video_Companionad:
					goto handle_Companionad

				case ffj_t_Video_Api:
					goto handle_Api

				case ffj_t_Video_Companiontype:
					goto handle_Companiontype

				case ffj_t_Video_Ext:
					goto handle_Ext

				case ffj_t_Videono_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Mimes:

	/* handler: uj.Mimes type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Mimes = nil
		} else {

			uj.Mimes = make([]string, 0)

			wantVal := true

			for {

				var v string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						v = string(string(outBuf))

					}
				}

				uj.Mimes = append(uj.Mimes, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Minduration:

	/* handler: uj.Minduration type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Minduration = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Maxduration:

	/* handler: uj.Maxduration type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Maxduration = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Protocol:

	/* handler: uj.Protocol type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Protocol = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Protocols:

	/* handler: uj.Protocols type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Protocols = nil
		} else {

			uj.Protocols = make([]int, 0)

			wantVal := true

			for {

				var v int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						v = int(tval)

					}
				}

				uj.Protocols = append(uj.Protocols, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_W:

	/* handler: uj.W type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.W = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_H:

	/* handler: uj.H type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.H = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Startdelay:

	/* handler: uj.Startdelay type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Startdelay = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Linearity:

	/* handler: uj.Linearity type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Linearity = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Sequence:

	/* handler: uj.Sequence type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Sequence = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Battr:

	/* handler: uj.Battr type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Battr = nil
		} else {

			uj.Battr = make([]int, 0)

			wantVal := true

			for {

				var v int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						v = int(tval)

					}
				}

				uj.Battr = append(uj.Battr, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Maxextended:

	/* handler: uj.Maxextended type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Maxextended = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Minbitrate:

	/* handler: uj.Minbitrate type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Minbitrate = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Maxbitrate:

	/* handler: uj.Maxbitrate type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Maxbitrate = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Boxingallowed:

	/* handler: uj.Boxingallowed type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Boxingallowed = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Playbackmethod:

	/* handler: uj.Playbackmethod type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Playbackmethod = nil
		} else {

			uj.Playbackmethod = make([]int, 0)

			wantVal := true

			for {

				var v int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						v = int(tval)

					}
				}

				uj.Playbackmethod = append(uj.Playbackmethod, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Delivery:

	/* handler: uj.Delivery type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Delivery = nil
		} else {

			uj.Delivery = make([]int, 0)

			wantVal := true

			for {

				var v int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						v = int(tval)

					}
				}

				uj.Delivery = append(uj.Delivery, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Pos:

	/* handler: uj.Pos type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Pos = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Companionad:

	/* handler: uj.Companionad type=[]openrtb.Banner kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Companionad = nil
		} else {

			uj.Companionad = make([]Banner, 0)

			wantVal := true

			for {

				var v Banner

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=openrtb.Banner kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

						state = fflib.FFParse_after_value
						goto mainparse
					}

					err = v.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.Companionad = append(uj.Companionad, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Api:

	/* handler: uj.Api type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Api = nil
		} else {

			uj.Api = make([]int, 0)

			wantVal := true

			for {

				var v int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						v = int(tval)

					}
				}

				uj.Api = append(uj.Api, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Companiontype:

	/* handler: uj.Companiontype type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Companiontype = nil
		} else {

			uj.Companiontype = make([]int, 0)

			wantVal := true

			for {

				var v int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						v = int(tval)

					}
				}

				uj.Companiontype = append(uj.Companiontype, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ext:

	/* handler: uj.Ext type=openrtb.Extensions kind=map quoted=false*/

	{
		/* Falling back. type=openrtb.Extensions kind=map */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Ext)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
