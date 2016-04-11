// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: asset.go
// DO NOT EDIT!

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Asset) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Asset) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "id":`)
	fflib.FormatBits2(buf, uint64(mj.Id), 10, mj.Id < 0)
	buf.WriteByte(',')
	if mj.Required != 0 {
		buf.WriteString(`"required":`)
		fflib.FormatBits2(buf, uint64(mj.Required), 10, mj.Required < 0)
		buf.WriteByte(',')
	}
	if mj.Title != nil {
		if true {
			buf.WriteString(`"title":`)

			{
				err = mj.Title.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

			buf.WriteByte(',')
		}
	}
	if mj.Img != nil {
		if true {
			buf.WriteString(`"img":`)

			{
				err = mj.Img.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

			buf.WriteByte(',')
		}
	}
	if mj.Video != nil {
		if true {
			/* Struct fall back. type=openrtb.NativeVideo kind=struct */
			buf.WriteString(`"video":`)
			err = buf.Encode(mj.Video)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if mj.Data != nil {
		if true {
			buf.WriteString(`"data":`)

			{
				err = mj.Data.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}
			}

			buf.WriteByte(',')
		}
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
	ffj_t_Assetbase = iota
	ffj_t_Assetno_such_key

	ffj_t_Asset_Id

	ffj_t_Asset_Required

	ffj_t_Asset_Title

	ffj_t_Asset_Img

	ffj_t_Asset_Video

	ffj_t_Asset_Data

	ffj_t_Asset_Ext
)

var ffj_key_Asset_Id = []byte("id")

var ffj_key_Asset_Required = []byte("required")

var ffj_key_Asset_Title = []byte("title")

var ffj_key_Asset_Img = []byte("img")

var ffj_key_Asset_Video = []byte("video")

var ffj_key_Asset_Data = []byte("data")

var ffj_key_Asset_Ext = []byte("ext")

func (uj *Asset) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Asset) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Assetbase
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
				currentKey = ffj_t_Assetno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffj_key_Asset_Data, kn) {
						currentKey = ffj_t_Asset_Data
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Asset_Ext, kn) {
						currentKey = ffj_t_Asset_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_Asset_Id, kn) {
						currentKey = ffj_t_Asset_Id
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Asset_Img, kn) {
						currentKey = ffj_t_Asset_Img
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffj_key_Asset_Required, kn) {
						currentKey = ffj_t_Asset_Required
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Asset_Title, kn) {
						currentKey = ffj_t_Asset_Title
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffj_key_Asset_Video, kn) {
						currentKey = ffj_t_Asset_Video
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Asset_Ext, kn) {
					currentKey = ffj_t_Asset_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Asset_Data, kn) {
					currentKey = ffj_t_Asset_Data
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Asset_Video, kn) {
					currentKey = ffj_t_Asset_Video
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Asset_Img, kn) {
					currentKey = ffj_t_Asset_Img
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Asset_Title, kn) {
					currentKey = ffj_t_Asset_Title
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Asset_Required, kn) {
					currentKey = ffj_t_Asset_Required
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Asset_Id, kn) {
					currentKey = ffj_t_Asset_Id
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Assetno_such_key
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

				case ffj_t_Asset_Id:
					goto handle_Id

				case ffj_t_Asset_Required:
					goto handle_Required

				case ffj_t_Asset_Title:
					goto handle_Title

				case ffj_t_Asset_Img:
					goto handle_Img

				case ffj_t_Asset_Video:
					goto handle_Video

				case ffj_t_Asset_Data:
					goto handle_Data

				case ffj_t_Asset_Ext:
					goto handle_Ext

				case ffj_t_Assetno_such_key:
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

handle_Id:

	/* handler: uj.Id type=int kind=int quoted=false*/

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

			uj.Id = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Required:

	/* handler: uj.Required type=int kind=int quoted=false*/

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

			uj.Required = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Title:

	/* handler: uj.Title type=openrtb.Title kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Title = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Title == nil {
			uj.Title = new(Title)
		}

		err = uj.Title.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Img:

	/* handler: uj.Img type=openrtb.Image kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Img = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Img == nil {
			uj.Img = new(Image)
		}

		err = uj.Img.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Video:

	/* handler: uj.Video type=openrtb.NativeVideo kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.NativeVideo kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Video)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Data:

	/* handler: uj.Data type=openrtb.NativeData kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Data = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Data == nil {
			uj.Data = new(NativeData)
		}

		err = uj.Data.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
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

func (mj *Image) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Image) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if mj.Type != 0 {
		buf.WriteString(`"type":`)
		fflib.FormatBits2(buf, uint64(mj.Type), 10, mj.Type < 0)
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
	if mj.Wmax != 0 {
		buf.WriteString(`"wmax":`)
		fflib.FormatBits2(buf, uint64(mj.Wmax), 10, mj.Wmax < 0)
		buf.WriteByte(',')
	}
	if mj.Hmax != 0 {
		buf.WriteString(`"hmax":`)
		fflib.FormatBits2(buf, uint64(mj.Hmax), 10, mj.Hmax < 0)
		buf.WriteByte(',')
	}
	if mj.Wmin != 0 {
		buf.WriteString(`"wmin":`)
		fflib.FormatBits2(buf, uint64(mj.Wmin), 10, mj.Wmin < 0)
		buf.WriteByte(',')
	}
	if mj.Hmin != 0 {
		buf.WriteString(`"hmin":`)
		fflib.FormatBits2(buf, uint64(mj.Hmin), 10, mj.Hmin < 0)
		buf.WriteByte(',')
	}
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
	ffj_t_Imagebase = iota
	ffj_t_Imageno_such_key

	ffj_t_Image_Type

	ffj_t_Image_W

	ffj_t_Image_H

	ffj_t_Image_Wmax

	ffj_t_Image_Hmax

	ffj_t_Image_Wmin

	ffj_t_Image_Hmin

	ffj_t_Image_Mimes

	ffj_t_Image_Ext
)

var ffj_key_Image_Type = []byte("type")

var ffj_key_Image_W = []byte("w")

var ffj_key_Image_H = []byte("h")

var ffj_key_Image_Wmax = []byte("wmax")

var ffj_key_Image_Hmax = []byte("hmax")

var ffj_key_Image_Wmin = []byte("wmin")

var ffj_key_Image_Hmin = []byte("hmin")

var ffj_key_Image_Mimes = []byte("mimes")

var ffj_key_Image_Ext = []byte("ext")

func (uj *Image) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Image) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Imagebase
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
				currentKey = ffj_t_Imageno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffj_key_Image_Ext, kn) {
						currentKey = ffj_t_Image_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffj_key_Image_H, kn) {
						currentKey = ffj_t_Image_H
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Image_Hmax, kn) {
						currentKey = ffj_t_Image_Hmax
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Image_Hmin, kn) {
						currentKey = ffj_t_Image_Hmin
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_Image_Mimes, kn) {
						currentKey = ffj_t_Image_Mimes
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Image_Type, kn) {
						currentKey = ffj_t_Image_Type
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffj_key_Image_W, kn) {
						currentKey = ffj_t_Image_W
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Image_Wmax, kn) {
						currentKey = ffj_t_Image_Wmax
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Image_Wmin, kn) {
						currentKey = ffj_t_Image_Wmin
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Image_Ext, kn) {
					currentKey = ffj_t_Image_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Image_Mimes, kn) {
					currentKey = ffj_t_Image_Mimes
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Image_Hmin, kn) {
					currentKey = ffj_t_Image_Hmin
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Image_Wmin, kn) {
					currentKey = ffj_t_Image_Wmin
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Image_Hmax, kn) {
					currentKey = ffj_t_Image_Hmax
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Image_Wmax, kn) {
					currentKey = ffj_t_Image_Wmax
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Image_H, kn) {
					currentKey = ffj_t_Image_H
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Image_W, kn) {
					currentKey = ffj_t_Image_W
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Image_Type, kn) {
					currentKey = ffj_t_Image_Type
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Imageno_such_key
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

				case ffj_t_Image_Type:
					goto handle_Type

				case ffj_t_Image_W:
					goto handle_W

				case ffj_t_Image_H:
					goto handle_H

				case ffj_t_Image_Wmax:
					goto handle_Wmax

				case ffj_t_Image_Hmax:
					goto handle_Hmax

				case ffj_t_Image_Wmin:
					goto handle_Wmin

				case ffj_t_Image_Hmin:
					goto handle_Hmin

				case ffj_t_Image_Mimes:
					goto handle_Mimes

				case ffj_t_Image_Ext:
					goto handle_Ext

				case ffj_t_Imageno_such_key:
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

handle_Type:

	/* handler: uj.Type type=int kind=int quoted=false*/

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

			uj.Type = int(tval)

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

handle_Wmax:

	/* handler: uj.Wmax type=int kind=int quoted=false*/

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

			uj.Wmax = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Hmax:

	/* handler: uj.Hmax type=int kind=int quoted=false*/

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

			uj.Hmax = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Wmin:

	/* handler: uj.Wmin type=int kind=int quoted=false*/

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

			uj.Wmin = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Hmin:

	/* handler: uj.Hmin type=int kind=int quoted=false*/

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

			uj.Hmin = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

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

func (mj *NativeData) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *NativeData) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "type":`)
	fflib.FormatBits2(buf, uint64(mj.Type), 10, mj.Type < 0)
	buf.WriteByte(',')
	if mj.Len != 0 {
		buf.WriteString(`"len":`)
		fflib.FormatBits2(buf, uint64(mj.Len), 10, mj.Len < 0)
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
	ffj_t_NativeDatabase = iota
	ffj_t_NativeDatano_such_key

	ffj_t_NativeData_Type

	ffj_t_NativeData_Len

	ffj_t_NativeData_Ext
)

var ffj_key_NativeData_Type = []byte("type")

var ffj_key_NativeData_Len = []byte("len")

var ffj_key_NativeData_Ext = []byte("ext")

func (uj *NativeData) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *NativeData) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_NativeDatabase
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
				currentKey = ffj_t_NativeDatano_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffj_key_NativeData_Ext, kn) {
						currentKey = ffj_t_NativeData_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffj_key_NativeData_Len, kn) {
						currentKey = ffj_t_NativeData_Len
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_NativeData_Type, kn) {
						currentKey = ffj_t_NativeData_Type
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_NativeData_Ext, kn) {
					currentKey = ffj_t_NativeData_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_NativeData_Len, kn) {
					currentKey = ffj_t_NativeData_Len
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_NativeData_Type, kn) {
					currentKey = ffj_t_NativeData_Type
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_NativeDatano_such_key
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

				case ffj_t_NativeData_Type:
					goto handle_Type

				case ffj_t_NativeData_Len:
					goto handle_Len

				case ffj_t_NativeData_Ext:
					goto handle_Ext

				case ffj_t_NativeDatano_such_key:
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

handle_Type:

	/* handler: uj.Type type=int kind=int quoted=false*/

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

			uj.Type = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Len:

	/* handler: uj.Len type=int kind=int quoted=false*/

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

			uj.Len = int(tval)

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

func (mj *Title) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Title) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "len":`)
	fflib.FormatBits2(buf, uint64(mj.Len), 10, mj.Len < 0)
	buf.WriteByte(',')
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
	ffj_t_Titlebase = iota
	ffj_t_Titleno_such_key

	ffj_t_Title_Len

	ffj_t_Title_Ext
)

var ffj_key_Title_Len = []byte("len")

var ffj_key_Title_Ext = []byte("ext")

func (uj *Title) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Title) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Titlebase
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
				currentKey = ffj_t_Titleno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffj_key_Title_Ext, kn) {
						currentKey = ffj_t_Title_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffj_key_Title_Len, kn) {
						currentKey = ffj_t_Title_Len
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Title_Ext, kn) {
					currentKey = ffj_t_Title_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Title_Len, kn) {
					currentKey = ffj_t_Title_Len
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Titleno_such_key
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

				case ffj_t_Title_Len:
					goto handle_Len

				case ffj_t_Title_Ext:
					goto handle_Ext

				case ffj_t_Titleno_such_key:
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

handle_Len:

	/* handler: uj.Len type=int kind=int quoted=false*/

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

			uj.Len = int(tval)

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
