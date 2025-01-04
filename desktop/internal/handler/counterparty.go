package handler

import (
	"desktop-templater-docx/internal/domain/entity"
	"desktop-templater-docx/internal/utils"
	"fmt"

	"github.com/lukasjarosch/go-docx"
)

// TODO
// Отдельная абстракция для контрагента?
// ...
type CounterpartyHandler struct {
	counterparties entity.Counterparties
}

func NewCounterpartyHandler() *CounterpartyHandler {
	return &CounterpartyHandler{
		counterparties: entity.Counterparties{},
	}
}

func (ch *CounterpartyHandler) MakeReplace(id int) {
	replaceMap, err := ch.getCounterpartyPlaceholderMap(id)
	if err != nil {
		fmt.Printf("[Error]: %v\n", err)
		return
	}

	// TODO вынести в настройки
	var pathToTemplate = "./templates/type" + ch.counterparties[id].Contract_type + ".docx"
	var pathToSave = "./replaced/" + ch.counterparties[id].Inn + ".docx"

	err = utils.PlaceholderReplacer(pathToTemplate, pathToSave, replaceMap)
	if err != nil {
		fmt.Printf("[Error]: %v\n", err)
		return
	}
}

func (ch *CounterpartyHandler) SetCounterparties(counterparties entity.Counterparties) {
	ch.counterparties = counterparties
}

func (ch *CounterpartyHandler) GetCounterparties() entity.Counterparties {
	return ch.counterparties
}

func (ch *CounterpartyHandler) GetCounterparty(pos int) (*entity.Counterparty, error) {
	if pos < 0 && pos > len(ch.counterparties) {
		return nil, fmt.Errorf("out of range")
	}

	return ch.counterparties[pos], nil
}

func (ch *CounterpartyHandler) getCounterpartyPlaceholderMap(pos int) (docx.PlaceholderMap, error) {
	if pos < 0 && pos > len(ch.counterparties) {
		return nil, fmt.Errorf("out of range")
	}

	placeholderMap := docx.PlaceholderMap{
		"A": ch.counterparties[pos].Code_ou,
		"B": ch.counterparties[pos].Inn,
		"C": ch.counterparties[pos].Institution_short_name,
		"D": ch.counterparties[pos].Institution_full_name,
		"E": ch.counterparties[pos].Address,
		"F": ch.counterparties[pos].City,
		"G": ch.counterparties[pos].Bank_details,
		"H": ch.counterparties[pos].Responsible_person_job_title,
		"I": ch.counterparties[pos].Responsible_person_short_name,
		"J": ch.counterparties[pos].Responsible_person_full_name,
		"K": ch.counterparties[pos].Responsible_person_full_name_genitive,
		"L": ch.counterparties[pos].Acting_on,
		"M": ch.counterparties[pos].Ikz_2025,
		"N": ch.counterparties[pos].Source_funding,
		"O": ch.counterparties[pos].Email,
		"P": ch.counterparties[pos].Phone_number,
		"Q": ch.counterparties[pos].Contract_form,
		"R": ch.counterparties[pos].Contract_type,
		"S": ch.counterparties[pos].Contract_number,
		"T": ch.counterparties[pos].Contract_formation_data,
		"U": ch.counterparties[pos].Responsible_person_job_title_genetive,
		"V": ch.counterparties[pos].Category,
	}

	return placeholderMap, nil
}
