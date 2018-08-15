/* 
 * Catapult REST API Reference
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.7.7
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package infrastructure

type AccountDto struct {

	Address string `json:"address"`

	AddressHeight *[]UInt64Dto `json:"addressHeight"`

	PublicKey string `json:"publicKey"`

	PublicKeyHeight *[]UInt64Dto `json:"publicKeyHeight"`

	Mosaics []MosaicDto `json:"mosaics"`

	Importance *[]UInt64Dto `json:"importance"`

	ImportanceHeight *[]UInt64Dto `json:"importanceHeight"`
}
