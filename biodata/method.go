package biodata

import (
	"fmt"
	"strings"
)

type ContractMethodBiodata interface {
	ListEmail() map[string][]string
	FilterEmail(name string) string
	GenerateUsername() string
	GenerateID() int
	GenerateAddress() string
	GenerateJob() string
	GenerateReason() string
}

/* LIST NAMES */
func ContractListEmail(contract ContractMethodBiodata) {
	fmt.Println("List Email ", contract.ListEmail())
}

func (b *Biodata) ListEmail() map[string][]string {
	arrEmail := []string{"Claude@gmail.com", "Leo@gmail.com", "Julian@gmail.com"}
	dataEmail := map[string][]string{
		"Emails": arrEmail,
	}
	return dataEmail
}

/* FILTER NAMES */
func ContractFilterEmail(contract ContractMethodBiodata, nama string) {
	fmt.Println("FIlter Email", contract.FilterEmail(nama))
}

func (b *Biodata) FilterEmail(name string) string {
	emails := b.ListEmail()
	for i := 0; i < len(emails); i++ {
		for j := 0; j < len(emails["Emails"]); j++ {
			if emails["Emails"][j] == name {
				b.Email = emails["Emails"][j]
			}
		}

	}
	return b.Email
}

/* GENERATE ID */

func ContractGenerateID(contract ContractMethodBiodata) {
	fmt.Println("Generate ID ", contract.GenerateID())
}

func (b *Biodata) GenerateID() int {
	emails := b.ListEmail()
	for i := 0; i < len(emails); i++ {
		for j := 0; j < len(emails["Emails"]); j++ {
			if emails["Emails"][j] == b.Email {
				b.ID = j
			}
		}
	}

	return b.ID
}

/* GENERATE USERNAME */

func ContractGenerateUsername(contract ContractMethodBiodata) {
	fmt.Println("Generate Username", contract.GenerateUsername())
}

func (b *Biodata) GenerateUsername() string {
	newUsername := strings.Split(b.Email, "@")
	b.Username = newUsername[0]
	return b.Username
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
	b.Reason = "Alasan " + b.Username
	return b.Reason
}
