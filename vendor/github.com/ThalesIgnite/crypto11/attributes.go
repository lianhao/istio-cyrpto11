package crypto11

import (
	"errors"
	"fmt"

	"github.com/miekg/pkcs11"
)

// AttributeType represents a PKCS#11 CK_ATTRIBUTE value.
type AttributeType = uint

// Attribute represents a PKCS#11 CK_ATTRIBUTE type.
type Attribute = pkcs11.Attribute

//noinspection GoUnusedConst
const (
	CkaClass                  = AttributeType(0x00000000)
	CkaToken                  = AttributeType(0x00000001)
	CkaPrivate                = AttributeType(0x00000002)
	CkaLabel                  = AttributeType(0x00000003)
	CkaApplication            = AttributeType(0x00000010)
	CkaValue                  = AttributeType(0x00000011)
	CkaObjectId               = AttributeType(0x00000012)
	CkaCertificateType        = AttributeType(0x00000080)
	CkaIssuer                 = AttributeType(0x00000081)
	CkaSerialNumber           = AttributeType(0x00000082)
	CkaAcIssuer               = AttributeType(0x00000083)
	CkaOwner                  = AttributeType(0x00000084)
	CkaAttrTypes              = AttributeType(0x00000085)
	CkaTrusted                = AttributeType(0x00000086)
	CkaCertificateCategory    = AttributeType(0x00000087)
	CkaJavaMIDPSecurityDomain = AttributeType(0x00000088)
	CkaUrl                    = AttributeType(0x00000089)
	CkaHashOfSubjectPublicKey = AttributeType(0x0000008A)
	CkaHashOfIssuerPublicKey  = AttributeType(0x0000008B)
	CkaNameHashAlgorithm      = AttributeType(0x0000008C)
	CkaCheckValue             = AttributeType(0x00000090)

	CkaKeyType         = AttributeType(0x00000100)
	CkaSubject         = AttributeType(0x00000101)
	CkaId              = AttributeType(0x00000102)
	CkaSensitive       = AttributeType(0x00000103)
	CkaEncrypt         = AttributeType(0x00000104)
	CkaDecrypt         = AttributeType(0x00000105)
	CkaWrap            = AttributeType(0x00000106)
	CkaUnwrap          = AttributeType(0x00000107)
	CkaSign            = AttributeType(0x00000108)
	CkaSignRecover     = AttributeType(0x00000109)
	CkaVerify          = AttributeType(0x0000010A)
	CkaVerifyRecover   = AttributeType(0x0000010B)
	CkaDerive          = AttributeType(0x0000010C)
	CkaStartDate       = AttributeType(0x00000110)
	CkaEndDate         = AttributeType(0x00000111)
	CkaModulus         = AttributeType(0x00000120)
	CkaModulusBits     = AttributeType(0x00000121)
	CkaPublicExponent  = AttributeType(0x00000122)
	CkaPrivateExponent = AttributeType(0x00000123)
	CkaPrime1          = AttributeType(0x00000124)
	CkaPrime2          = AttributeType(0x00000125)
	CkaExponent1       = AttributeType(0x00000126)
	CkaExponent2       = AttributeType(0x00000127)
	CkaCoefficient     = AttributeType(0x00000128)
	CkaPublicKeyInfo   = AttributeType(0x00000129)
	CkaPrime           = AttributeType(0x00000130)
	CkaSubprime        = AttributeType(0x00000131)
	CkaBase            = AttributeType(0x00000132)

	CkaPrimeBits    = AttributeType(0x00000133)
	CkaSubprimeBits = AttributeType(0x00000134)
	/* (To retain backwards-compatibility) */
	CkaSubPrimeBits = CkaSubprimeBits

	CkaValueBits        = AttributeType(0x00000160)
	CkaValueLen         = AttributeType(0x00000161)
	CkaExtractable      = AttributeType(0x00000162)
	CkaLocal            = AttributeType(0x00000163)
	CkaNeverExtractable = AttributeType(0x00000164)
	CkaAlwaysSensitive  = AttributeType(0x00000165)
	CkaKeyGenMechanism  = AttributeType(0x00000166)

	CkaModifiable = AttributeType(0x00000170)
	CkaCopyable   = AttributeType(0x00000171)

	/* new for v2.40 */
	CkaDestroyable = AttributeType(0x00000172)

	/* CKA_ECDSA_PARAMS is deprecated in v2.11,
	 * CKA_EC_PARAMS is preferred. */
	CkaEcdsaParams = AttributeType(0x00000180)
	CkaEcParams    = AttributeType(0x00000180)

	CkaEcPoint = AttributeType(0x00000181)

	/* CKA_SECONDARY_AUTH, CKA_AUTH_PIN_FLAGS,
	 * are new for v2.10. Deprecated in v2.11 and onwards. */
	CkaSecondaryAuth = AttributeType(0x00000200) /* Deprecated */
	CkaAuthPinFlags  = AttributeType(0x00000201) /* Deprecated */

	CkaAlwaysAuthenticate = AttributeType(0x00000202)

	CkaWrapWithTrusted = AttributeType(0x00000210)

	ckfArrayAttribute = AttributeType(0x40000000)

	CkaWrapTemplate   = ckfArrayAttribute | AttributeType(0x00000211)
	CkaUnwrapTemplate = ckfArrayAttribute | AttributeType(0x00000212)

	CkaOtpFormat               = AttributeType(0x00000220)
	CkaOtpLength               = AttributeType(0x00000221)
	CkaOtpTimeInterval         = AttributeType(0x00000222)
	CkaOtpUserFriendlyMode     = AttributeType(0x00000223)
	CkaOtpChallengeRequirement = AttributeType(0x00000224)
	CkaOtpTimeRequirement      = AttributeType(0x00000225)
	CkaOtpCounterRequirement   = AttributeType(0x00000226)
	CkaOtpPinRequirement       = AttributeType(0x00000227)
	CkaOtpCounter              = AttributeType(0x0000022E)
	CkaOtpTime                 = AttributeType(0x0000022F)
	CkaOtpUserIdentifier       = AttributeType(0x0000022A)
	CkaOtpServiceIdentifier    = AttributeType(0x0000022B)
	CkaOtpServiceLogo          = AttributeType(0x0000022C)
	CkaOtpServiceLogoType      = AttributeType(0x0000022D)

	CkaGOSTR3410Params = AttributeType(0x00000250)
	CkaGOSTR3411Params = AttributeType(0x00000251)
	CkaGOST28147Params = AttributeType(0x00000252)

	CkaHwFeatureType = AttributeType(0x00000300)
	CkaResetOnInit   = AttributeType(0x00000301)
	CkaHasReset      = AttributeType(0x00000302)

	CkaPixelX                 = AttributeType(0x00000400)
	CkaPixelY                 = AttributeType(0x00000401)
	CkaResolution             = AttributeType(0x00000402)
	CkaCharRows               = AttributeType(0x00000403)
	CkaCharColumns            = AttributeType(0x00000404)
	CkaColor                  = AttributeType(0x00000405)
	CkaBitsPerPixel           = AttributeType(0x00000406)
	CkaCharSets               = AttributeType(0x00000480)
	CkaEncodingMethods        = AttributeType(0x00000481)
	CkaMimeTypes              = AttributeType(0x00000482)
	CkaMechanismType          = AttributeType(0x00000500)
	CkaRequiredCmsAttributes  = AttributeType(0x00000501)
	CkaDefaultCmsAttributes   = AttributeType(0x00000502)
	CkaSupportedCmsAttributes = AttributeType(0x00000503)
	CkaAllowedMechanisms      = ckfArrayAttribute | AttributeType(0x00000600)
)

// NewAttribute is a helper function that populates a new Attribute for common data types. This function will
// return an error if value is not of type bool, int, uint, string, []byte or time.Time (or is nil).
func NewAttribute(attributeType AttributeType, value interface{}) (a *Attribute, err error) {
	// catch any panics from the pkcs11.NewAttribute() call to handle the error cleanly
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("failed creating Attribute: %v", r))
		}
	}()

	pAttr := pkcs11.NewAttribute(attributeType, value)
	return pAttr, nil
}

// CopyAttribute returns a deep copy of the given Attribute.
func CopyAttribute(a *Attribute) *Attribute {
	var value []byte
	if a.Value != nil && len(a.Value) > 0 {
		value = append([]byte(nil), a.Value...)
	}
	return &pkcs11.Attribute{
		Type:  a.Type,
		Value: value,
	}
}

// An AttributeSet groups together operations that are common for a collection of Attributes.
type AttributeSet map[AttributeType]*Attribute

// NewAttributeSet creates an empty AttributeSet.
func NewAttributeSet() AttributeSet {
	return make(AttributeSet)
}

// Set stores a new attribute in the AttributeSet. Any existing value will be overwritten. This function will return an
// error if value is not of type bool, int, uint, string, []byte or time.Time (or is nil).
func (a AttributeSet) Set(attributeType AttributeType, value interface{}) error {
	attr, err := NewAttribute(attributeType, value)
	if err != nil {
		return err
	}
	a[attributeType] = attr
	return nil
}

// AddIfNotPresent adds the attributes if the Attribute Type is not already present in the AttributeSet.
func (a AttributeSet) AddIfNotPresent(additional []*Attribute) {
	for _, additionalAttr := range additional {
		// Only add the attribute if it is not already present in the Attribute map
		if _, ok := a[additionalAttr.Type]; !ok {
			a[additionalAttr.Type] = additionalAttr
		}
	}
}

// ToSlice returns a deep copy of Attributes contained in the AttributeSet.
func (a AttributeSet) ToSlice() []*Attribute {
	var attributes []*Attribute
	for _, v := range a {
		duplicateAttr := CopyAttribute(v)
		attributes = append(attributes, duplicateAttr)
	}
	return attributes
}

// Copy returns a deep copy of the AttributeSet. This function will return an error if value is not of type
// bool, int, uint, string, []byte or time.Time (or is nil).
func (a AttributeSet) Copy() AttributeSet {
	b := NewAttributeSet()
	for _, v := range a {
		b[v.Type] = CopyAttribute(v)
	}
	return b
}

// NewAttributeSetWithID is a helper function that populates a new slice of Attributes with the provided ID.
// This function returns an error if the ID is an empty slice.
func NewAttributeSetWithID(id []byte) (AttributeSet, error) {
	if err := notNilBytes(id, "id"); err != nil {
		return nil, err
	}
	a := NewAttributeSet()
	_ = a.Set(CkaId, id) // error not possible for []byte
	return a, nil
}

// NewAttributeSetWithIDAndLabel is a helper function that populates a new slice of Attributes with the
// provided ID and Label. This function returns an error if either the ID or the Label is an empty slice.
func NewAttributeSetWithIDAndLabel(id, label []byte) (a AttributeSet, err error) {
	if a, err = NewAttributeSetWithID(id); err != nil {
		return nil, err
	}

	if err := notNilBytes(label, "label"); err != nil {
		return nil, err
	}

	_ = a.Set(CkaLabel, label) // error not possible with []byte
	return a, nil
}
