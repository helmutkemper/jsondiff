package diffServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
)

type Data struct {
	Data      []Pessoa `json:"data"`
	logText   string
	elementA  []any
	elementB  []any
	errorFunc func(error)
}

func (e *Data) SetErrorFunc(f func(error)) {
	e.errorFunc = f
}

func (e *Data) GetLog() string {
	return e.logText
}

func (e *Data) GetElementA() []any {
	return e.elementA
}

func (e *Data) GetElementB() []any {
	return e.elementB
}

func (e *Data) GetTotalElements() int {
	return len(e.elementA)
}

func (e *Data) GetElements(index int) (elementA, elementB []byte) {
	var err error
	if index >= 0 && index < len(e.elementA) {
		elementA, err = json.MarshalIndent(e.elementA[index], "", "  ")
		if err != nil {
			if e.errorFunc != nil {
				e.errorFunc(errors.Join(errors.New("Data().GetElements().json.MarshalIndent(A).error"), err))
			}
			return
		}

		elementB, err = json.MarshalIndent(e.elementB[index], "", "  ")
		if err != nil {
			if e.errorFunc != nil {
				e.errorFunc(errors.Join(errors.New("Data().GetElements().json.MarshalIndent(B).error"), err))
			}
			return
		}

		return elementA, elementB
	}

	return nil, nil
}

type Endereco struct {
	Rua    string `json:"rua"`
	Cidade string `json:"cidade"`
	Bairro string `json:"bairro"`
	CEP    string `json:"cep"`
	Numero int    `json:"numero"`
}

type Pessoa struct {
	ID        int      `json:"id"`
	Cpf       int      `json:"cpf"`
	Nome      string   `json:"nome"`
	Sobrenome string   `json:"sobrenome"`
	Idade     int      `json:"idade"`
	Endereco  Endereco `json:"endereco"`
}

// Função para criar uma chave única a partir de `n` chaves específicas
func (e *Data) makeUniqueKeys(v any, campos []string) string {
	val := reflect.ValueOf(v)
	tipo := reflect.TypeOf(v)

	var chave string
	for _, campo := range campos {
		field := val.FieldByName(campo)
		if field.IsValid() {
			chave += fmt.Sprintf("%v|", field.Interface()) // Concatena os valores das chaves
		} else {
			//e.logText += fmt.Sprintf("Campo %s não encontrado na struct %s", campo, tipo.Name())
			if e.errorFunc != nil {
				e.errorFunc(errors.Join(errors.New("Data().makeUniqueKeys().error"), errors.New(fmt.Sprintf("Campo %s não encontrado na struct %s", campo, tipo.Name()))))
			}
			return ""

		}
	}
	return chave
}

func (e *Data) Compare(data any, keys []string) {
	e.elementA = make([]any, 0)
	e.elementB = make([]any, 0)

	valA := reflect.ValueOf(e.Data)
	valB := reflect.ValueOf(data)

	if valA.Kind() != reflect.Slice || valB.Kind() != reflect.Slice {
		if e.errorFunc != nil {
			e.errorFunc(errors.Join(errors.New("Data().makeUniqueKeys().error"), errors.New("ss parâmetros devem ser slices")))
		}
		return
	}

	// Criar um mapa de busca para os elementos de B
	mapA := make(map[string]interface{})
	for i := 0; i < valA.Len(); i++ {
		elem := valA.Index(i).Interface()
		chave := e.makeUniqueKeys(elem, keys)
		mapA[chave] = elem
	}
	mapB := make(map[string]interface{})
	for i := 0; i < valB.Len(); i++ {
		elem := valB.Index(i).Interface()
		chave := e.makeUniqueKeys(elem, keys)
		mapB[chave] = elem
	}

	// Percorrer A e buscar correspondentes em B
	for i := 0; i < valA.Len(); i++ {
		elemA := valA.Index(i).Interface()
		chave := e.makeUniqueKeys(elemA, keys)

		if elemB, existe := mapB[chave]; existe {
			// Comparar os structs inteiros
			if !reflect.DeepEqual(elemA, elemB) {
				//e.logText += fmt.Sprintf("Diferença encontrada na chave %s:\n A: %+v\n B: %+v\n\n", chave, elemA, elemB)
				e.elementA = append(e.elementA, elemA)
				e.elementB = append(e.elementB, elemB)
			}
		} else {
			e.logText += fmt.Sprintf("Elemento de A com chave %s não encontrado em B\n", chave)
		}
	}

	//----------------------------------------------------------------------------------

	for i := 0; i < valB.Len(); i++ {
		elemB := valB.Index(i).Interface()
		chave := e.makeUniqueKeys(elemB, keys)
		if _, existe := mapA[chave]; !existe {
			e.logText += fmt.Sprintf("Elemento de B com chave %s não encontrado em A\n", chave)
		}
	}

}

// Shuffle Altera a ordem dos structs dentro do dado sem alterar o dado propriamente dito
func (e *Data) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(e.Data), func(i, j int) {
		e.Data[i], e.Data[j] = e.Data[j], e.Data[i]
	})
}

// CopyAndPrepare Copia o dado original e o altera n chaves (numberOfKeys) dentro de um struct do dado para x dados (interactions)
func (e *Data) CopyAndPrepare(data Data, interactions, numberOfKeys, deleteKeys int) {
	e.copy(data)
	e.blurring(interactions, numberOfKeys)
	e.shuffle()

	e.deleteKeys(data, deleteKeys)
	e.deleteKeys(*e, deleteKeys)
}

func (e *Data) deleteKeys(data Data, deleteKeys int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < deleteKeys; i++ {
		index := rand.Intn(len(data.Data))
		data.Data = append(data.Data[:index], data.Data[index+1:]...)
	}
}

// copy o dado original
//
//	O sistema necessita de dados quase iguais, por isto a cópia
func (e *Data) copy(data Data) {
	e.Data = make([]Pessoa, len(data.Data))
	copy(e.Data, data.Data)
}

// blurringFunctions Monta a lista de funções de embaraçamento
func (e *Data) blurringFunctions() (list []func(pessoa Pessoa) (embarrassedPerson Pessoa)) {
	list = make([]func(pessoa Pessoa) (embarrassedPerson Pessoa), 0)
	list = append(
		list,
		func(pessoa Pessoa) (embarrassedPerson Pessoa) {
			pessoa.Nome = gofakeit.FirstName()
			return pessoa
		})
	list = append(
		list,
		func(pessoa Pessoa) (embarrassedPerson Pessoa) {
			pessoa.Sobrenome = gofakeit.LastName()
			return pessoa
		})
	list = append(
		list,
		func(pessoa Pessoa) (embarrassedPerson Pessoa) {
			pessoa.Idade = gofakeit.Number(18, 99)
			return pessoa
		})
	list = append(
		list,
		func(pessoa Pessoa) (embarrassedPerson Pessoa) {
			pessoa.Endereco.Rua = gofakeit.Street()
			return pessoa
		})
	list = append(
		list,
		func(pessoa Pessoa) (embarrassedPerson Pessoa) {
			pessoa.Endereco.Cidade = gofakeit.City()
			return pessoa
		})
	list = append(
		list,
		func(pessoa Pessoa) (embarrassedPerson Pessoa) {
			pessoa.Endereco.Bairro = gofakeit.Word()
			return pessoa
		})
	list = append(
		list,
		func(pessoa Pessoa) (embarrassedPerson Pessoa) {
			pessoa.Endereco.CEP = gofakeit.Zip()
			return pessoa
		})
	list = append(
		list,
		func(pessoa Pessoa) (embarrassedPerson Pessoa) {
			pessoa.Endereco.Numero = gofakeit.Number(1, 1000)
			return pessoa
		})
	return
}

// blurring Altera n chaves (numberOfKeys) dentro de um struct do dado para x dados (interactions)
//
//	Exemplo, escolhe um struct de forma aleatória e altera a chave pessoa
func (e *Data) blurring(interactions, numberOfKeys int) {
	rand.Seed(time.Now().UnixNano())

	// Monta a lista de funções de embaraçamento
	blurringFunc := e.blurringFunctions()

	// Define as quantidades máximas para a escolha aleatória
	maxBlurringFunc := len(blurringFunc)
	maxData := len(e.Data)

	// Escolhe um dado para embaraçar
	for i := 0; i < interactions; i++ {
		index := rand.Intn(maxData)

		// Escolhe uma chave para embaraçar
		for j := 0; j < numberOfKeys; j++ {
			indexFunc := rand.Intn(maxBlurringFunc)

			pessoa := e.Data[index]
			pessoa = blurringFunc[indexFunc](pessoa)
			e.Data[index] = pessoa
		}
	}
}

func (e *Data) Get() []Pessoa {
	return e.Data
}

func (e *Data) Init(amountOfData int) {

	e.elementA = make([]any, 0)
	e.elementB = make([]any, 0)

	e.Data = make([]Pessoa, 0, 0)

	for i := 1; i <= amountOfData; i++ {
		pessoa := Pessoa{
			ID:        i,
			Cpf:       gofakeit.Number(1, 1000),
			Nome:      gofakeit.FirstName(),
			Sobrenome: gofakeit.LastName(),
			Idade:     gofakeit.Number(18, 99),
			Endereco: Endereco{
				Rua:    gofakeit.Street(),
				Cidade: gofakeit.City(),
				Bairro: gofakeit.Word(),
				CEP:    gofakeit.Zip(),
				Numero: gofakeit.Number(1, 1000),
			},
		}
		e.Data = append(e.Data, pessoa)
	}

}

func (e *Data) ___Sort(original []byte) (js []byte, err error) {
	data := make([]map[string]any, 0)
	err = json.Unmarshal(original, &data)
	if err != nil {
		panic(err)
	}

	e.___findAndSortArrays(data, "data", "id")
	js, err = json.MarshalIndent(&data, "", "  ")
	if err != nil {
		panic(err)
	}

	return
}

// Função para ordenar um slice de map[string]any por uma chave específica
func (e *Data) _sortSliceByKey(slice []map[string]any, key string) {
	sort.Slice(slice, func(i, j int) bool {
		// Usando reflect para obter os valores dinamicamente
		valI := reflect.ValueOf(slice[i][key])
		valJ := reflect.ValueOf(slice[j][key])

		// Comparando os valores com base no tipo
		switch valI.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return valI.Int() < valJ.Int()
		case reflect.Float32, reflect.Float64:
			return valI.Float() < valJ.Float()
		case reflect.String:
			return valI.String() < valJ.String()
		default:
			// Caso o tipo não seja suportado, não ordena
			return false
		}
	})
}

// Função para encontrar e ordenar um array específico no JSON
func (e *Data) ___findAndSortArrays(data any, path, key string) {
	// Divide o caminho em partes (ex: "nested.data" -> ["nested", "data"])
	parts := strings.Split(path, ".")
	if len(parts) == 0 {
		return
	}

	// Navega até o array desejado
	current := data
	for i, part := range parts {
		if reflect.TypeOf(current).Kind() == reflect.Map {
			mapData := current.(map[string]any)
			if val, exists := mapData[part]; exists {
				if i == len(parts)-1 {
					// Se for a última parte do caminho, verifica se é um array de objetos
					if reflect.TypeOf(val).Kind() == reflect.Slice {
						slice := val.([]any)
						if len(slice) > 0 && reflect.TypeOf(slice[0]).Kind() == reflect.Map {
							// Converte para []map[string]any e ordena
							typedSlice := make([]map[string]any, len(slice))
							for i, item := range slice {
								typedSlice[i] = item.(map[string]any)
							}
							e._sortSliceByKey(typedSlice, key)
							// Atualiza o valor no mapa original
							mapData[part] = typedSlice
						}
					}
				} else {
					// Continua navegando
					current = val
				}
			} else {
				// Chave não encontrada
				return
			}
		} else {
			// Tipo inválido
			return
		}
	}
}
