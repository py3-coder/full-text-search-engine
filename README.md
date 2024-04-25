# full-text-search-engine
Go Full Text Search Engine

## Overview
This project is a Full-Text Search Engine implemented in Go using the inverted index method. It allows users to efficiently search through a collection of documents by indexing the words in the documents and then quickly finding the documents that contain a particular word or set of words.

## Features
- **Inverted Index:** Utilizes the inverted index data structure to efficiently map words to the documents they appear in.
- **Query Processing:** Supports querying for single words, phrases, and boolean expressions.
- **Scalable:** Designed to handle large collections of documents efficiently.
- **Concurrency:** Utilizes Go's concurrency features to handle multiple search queries simultaneously.
- **Easy to Use:** Provides a simple interface for indexing documents and performing searches.

## Installation
To install the Full-Text Search Engine, you need to have Go installed on your system. Once you have Go installed, you can clone the repository and build the project using the following commands:

git clone <repository_url>
cd Full-Text-Search-Engine
go build

## Usage
1. **Indexing Documents:** Before performing searches, you need to index the documents.
2. **Searching:** Once the documents are indexed. You can perform Searching.

## Steps Involved:
1. **Tokenisation** - Converting the Search String in Tokens
2. **Filtering** -  StopwordFilter , StemmerFilter
3. **Stemmer** - We used Porter Stemmer Algo. - [READ](https://vijinimallawaarachchi.com/2017/05/09/porter-stemming-algorithm/)
4. **Inverted Index** - Created a inverted index map.
5. **Search**  - Inverted Index will be Search in Index Token with Intersection.

   

