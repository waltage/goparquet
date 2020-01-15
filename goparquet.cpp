
#include "goparquet.h"

using namespace std;


parquet::WriterProperties::Builder PROPERTY_BUILDER;

shared_ptr<parquet::schema::ColumnPath> pq_column(const string dotpath) {
    return parquet::schema::ColumnPath::FromDotString(dotpath);
}

parquet::Encoding::type pq_encoding(const string encoding) {

    if (encoding == "PLAIN") return parquet::Encoding::PLAIN;
    if (encoding == "PLAIN_DICTIONARY") return parquet::Encoding::PLAIN_DICTIONARY;
    if (encoding == "RLE") return parquet::Encoding::RLE;
    if (encoding == "BIT_PACKED") return parquet::Encoding::BIT_PACKED;
    if (encoding == "DELTA_BINARY_PACKED") return parquet::Encoding::DELTA_BINARY_PACKED;
    if (encoding == "DELTA_LENGTH_BYTE_ARRAY") return parquet::Encoding::DELTA_LENGTH_BYTE_ARRAY;
    if (encoding == "DELTA_BYTE_ARRAY") return parquet::Encoding::DELTA_BYTE_ARRAY;
    if (encoding == "RLE_DICTIONARY") return parquet::Encoding::RLE_DICTIONARY;
    return parquet::Encoding::UNKNOWN;

}

parquet::Compression::type pq_compression(const string compression) {
    if (compression == "SNAPPY") return parquet::Compression::SNAPPY;
    if (compression == "GZIP") return parquet::Compression::GZIP;
    if (compression == "LZO") return parquet::Compression::LZO;
    if (compression == "BROTLI") return parquet::Compression::BROTLI;
    if (compression == "LZ4") return parquet::Compression::LZ4;
    if (compression == "ZSTD") return parquet::Compression::ZSTD;
    return parquet::Compression::UNCOMPRESSED;
}

shared_ptr<arrow::Table> arrow_file_to_table(const string path) {
    arrow::Status statusz;
    shared_ptr<arrow::io::ReadableFile> file;
    statusz = arrow::io::ReadableFile::Open(path, &file);


    shared_ptr<arrow::ipc::RecordBatchFileReader> reader;
    statusz = arrow::ipc::RecordBatchFileReader::Open(file.get(), &reader);

    vector<shared_ptr<arrow::RecordBatch>> batches;
    for (int i=0; i < reader->num_record_batches(); ++i) {
        shared_ptr<arrow::RecordBatch> chunk;
        statusz = reader->ReadRecordBatch(i, &chunk);

        batches.push_back(chunk);
    }

    shared_ptr<arrow::Table> table;
    statusz = arrow::Table::FromRecordBatches(reader->schema(), batches, &table);

    return table;
}

void arrow_table_to_parquet(shared_ptr<arrow::Table> table, const string outpath, long chunksize) {
    arrow::Status statusz;
    
    shared_ptr<arrow::io::FileOutputStream> outfile;
    shared_ptr<parquet::WriterProperties> properties = PROPERTY_BUILDER.build();

    statusz = arrow::io::FileOutputStream::Open(outpath, &outfile);
    statusz = parquet::arrow::WriteTable(*table, arrow::default_memory_pool(), outfile, chunksize, properties);


}

void SetCompression(const string compression) {
    PROPERTY_BUILDER.compression(pq_compression(compression));
}

void SetColumnCompression(const string dotpath, const string compression) {
    PROPERTY_BUILDER.compression(pq_column(dotpath), pq_compression(compression));
}

void SetColumnEncoding(const string dotpath, const string encoding) {
    PROPERTY_BUILDER.encoding(pq_column(dotpath), pq_encoding(encoding));
}

void EnableColumnDictionary(const string dotpath) {
    PROPERTY_BUILDER.enable_dictionary(pq_column(dotpath));
}

void DisableColumnDictionary(const string dotpath) {
    PROPERTY_BUILDER.disable_dictionary(pq_column(dotpath));
}

void ConvertArrowFileToParquet(const string arrowfile, const string parquetfile, long chunksize) {
    PROPERTY_BUILDER.created_by("goparquet arrow file converter");
    arrow_table_to_parquet(arrow_file_to_table(arrowfile), parquetfile, chunksize);
}




