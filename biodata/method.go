package biodata

import (
	"fmt"
)

type ContractMethodBiodata interface {
	ListNama() map[string][]string
	FilterNama(name string) string
	GenerateID() int
	GenerateAddress() string
	GenerateJob() string
	GenerateReason() string
}

/* LIST NAMES */
func ContractListNama(contract ContractMethodBiodata) {
	fmt.Println("List Nama ", contract.ListNama())
}

func (b *Biodata) ListNama() map[string][]string {
	arrNama := []string{"Claude", "Leo", "Julian"}
	dataNama := map[string][]string{
		"Names": arrNama,
	}
	return dataNama
}

/* FILTER NAMES */
func ContractFilterNama(contract ContractMethodBiodata, nama string) {
	fmt.Println("FIlter Nama", contract.FilterNama(nama))
}

func (b *Biodata) FilterNama(name string) string {
	names := b.ListNama()
	for i := 0; i < len(names); i++ {
		for j := 0; j < len(names["Names"]); j++ {
			if names["Names"][j] == name {
				b.Name = names["Names"][j]
			}
		}

	}
	return b.Name
}

/* GENERATE ID */

func ContractGenerateID(contract ContractMethodBiodata) {
	fmt.Println("Generate ID ", contract.GenerateID())
}

func (b *Biodata) GenerateID() int {
	names := b.ListNama()
	for i := 0; i < len(names); i++ {
		for j := 0; j < len(names["Names"]); j++ {
			if names["Names"][j] == b.Name {
				b.ID = j
			}
		}
	}

	return b.ID
}

/* GENERATE ADDRESS */

func ContractGenerateAddress(contract ContractMethodBiodata) {
	fmt.Println("Generate Address", contract.GenerateAddress())
}

func (b *Biodata) GenerateAddress() string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b.Address = "Kota " + string(alphabet[b.ID])

	return b.Address
}

/* GENERATE JOB */
func ContractGenerateJob(contract ContractMethodBiodata) {
	fmt.Println("Generate Job ", contract.GenerateJob())
}

func (b *Biodata) GenerateJob() string {
	b.Job = "Developer"
	return b.Job
}

/* GENERATE REASON */
func ContractGenerateReason(contract ContractMethodBiodata) {
	fmt.Println("Generate Reason ", contract.GenerateReason())
}
func (b *Biodata) GenerateReason() string {
	b.Reason = "Alasan " + b.Name
	return b.Reason
}
