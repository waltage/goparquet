#ifndef GOPARQUET_H
#define GOPARQUET_H

#include <arrow/api.h>
#include <arrow/io/api.h>
#include <arrow/ipc/api.h>
#include <arrow/memory_pool.h>
#include <parquet/arrow/writer.h>
#include <iostream>
#include <cstdlib>
#include <memory>

using namespace std;

shared_ptr<parquet::schema::ColumnPath> pq_column(const string dotpath);
parquet::Encoding::type pq_encoding(const string encoding);
parquet::Compression::type pq_compression(const string compression);
shared_ptr<arrow::Table> arrow_file_to_table(const string path);
void arrow_table_to_parquet(shared_ptr<arrow::Table> table, const string outpath, long chunksize);

/* Go Extern Methods */
void SetCompression(const string compression);
void SetColumnCompression(const string dotpath, const string compression);
void SetColumnEncoding(const string dotpath, const string encoding);
void EnableColumnDictionary(const string dotpath);
void DisableColumnDictionary(const string dotpath);
void ConvertArrowFileToParquet(const string arrowfile, const string parquetfile, long chunksize);

#endif
