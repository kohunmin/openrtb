package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NativeResponse", func() {
	var subject NativeResponse

	BeforeEach(func() {
		subject = NativeResponse{}
	})

	It("should have defaults", func() {
		subject.Assets = []ResponseAsset{ResponseAsset{}}
		subject.WithDefaults()

		Expect(subject.Ver).To(Equal("1"))
		Expect(subject.Assets[0].Required).To(Equal(0))
	})

	It("should parse native adm", func() {
		resp, err := ParseNativeAdmBytes(testFixtures.simpleNativeResponse)

		Expect(err).NotTo(HaveOccurred())
		nativeAdm := NativeAdm{
			&NativeResponse{
				Ver: "1",
				Assets: []ResponseAsset{
					ResponseAsset{
						Id:       1,
						Required: 0,
						Title:    &testFixtures.simpleTitle,
						Link:     &testFixtures.simpleLink,
					},
					ResponseAsset{
						Id:       2,
						Required: 0,
						Data:     &testFixtures.simpleData,
					},
					ResponseAsset{
						Id:       3,
						Required: 0,
						Img:      &testFixtures.simpleImg,
					},
					ResponseAsset{
						Id:       4,
						Required: 0,
						Data:     &testFixtures.installData,
						Link:     &testFixtures.simpleLink,
					},
				},
				Link:        &testFixtures.fullLink,
				Imptrackers: []string{"http: //a.com/a", "http: //b.com/b"},
			},
		}

		Expect(resp).To(Equal(&nativeAdm))
	})
})
