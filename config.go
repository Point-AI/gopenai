package gopenai

var (
	chatCompletionPathUrl = "/chat/completions"
	embeddingsPathUrl     = "/embeddings"
)

const (
	BaseURL = "https://api.openai.com/v1"
)

const (
	ModelGPT432K0613          = "gpt-4-32k-0613"
	ModelGPT432K0314          = "gpt-4-32k-0314"
	ModelGPT432K              = "gpt-4-32k"
	ModelGPT40613             = "gpt-4-0613"
	ModelGPT40314             = "gpt-4-0314"
	ModelGPT4                 = "gpt-4"
	ModelGPT3Dot5Turbo0613    = "gpt-3.5-turbo-0613"
	ModelGPT3Dot5Turbo0301    = "gpt-3.5-turbo-0301"
	ModelGPT3Dot5Turbo16K     = "gpt-3.5-turbo-16k"
	ModelGPT3Dot5Turbo16K0613 = "gpt-3.5-turbo-16k-0613"
	ModelGPT3Dot5Turbo        = "gpt-3.5-turbo"
)

const (
	ModelAdaSimilarity         = "text-similarity-ada-001"
	ModelBabbageSimilarity     = "text-similarity-babbage-001"
	ModelCurieSimilarity       = "text-similarity-curie-001"
	ModelDavinciSimilarity     = "text-similarity-davinci-001"
	ModelAdaSearchDocument     = "text-search-ada-doc-001"
	ModelAdaSearchQuery        = "text-search-ada-query-001"
	ModelBabbageSearchDocument = "text-search-babbage-doc-001"
	ModelBabbageSearchQuery    = "text-search-babbage-query-001"
	ModelCurieSearchDocument   = "text-search-curie-doc-001"
	ModelCurieSearchQuery      = "text-search-curie-query-001"
	ModelDavinciSearchDocument = "text-search-davinci-doc-001"
	ModelDavinciSearchQuery    = "text-search-davinci-query-001"
	ModelAdaCodeSearchCode     = "code-search-ada-code-001"
	ModelAdaCodeSearchText     = "code-search-ada-text-001"
	ModelBabbageCodeSearchCode = "code-search-babbage-code-001"
	ModelBabbageCodeSearchText = "code-search-babbage-text-001"
	ModelAdaEmbeddingV2        = "text-embedding-ada-002"
)
