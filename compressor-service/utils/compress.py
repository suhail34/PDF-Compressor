import fitz  # PyMuPDF
from connections.mongoConnections import MongoConnection
from gridfs import GridFSBucket
import logging
import tempfile
from bson import ObjectId
from io import BytesIO


def compress_pdf(name, file_ID):
    # Open the PDF file
    mongo_connection = MongoConnection()
    mongo_client = mongo_connection.get_client()
    db = mongo_client.myFiles
    fs = GridFSBucket(db)

    out = fs.open_download_stream(ObjectId(file_ID)).read()

    try:
        original_pdf = BytesIO(bytes(out))
        compressed_pdf_path = tempfile.gettempdir() + \
            f"/{name[:-4]}_compressed.pdf"

        # Open the original PDF
        with fitz.open("pdf", original_pdf) as pdf_document:
            # Create a new PDF with compressed content
            pdf_compressed = fitz.open()

            # Iterate through pages and add them to the new PDF
            for page_number in range(pdf_document.page_count):
                page = pdf_document[page_number]
                pix = page.get_pixmap()
                pdf_compressed.insert_page(
                    page_number,
                    width=pix.width,
                    height=pix.height
                )
                pdf_compressed[page_number].insert_text(
                    (0, 0), page.get_text())

            # Save the compressed PDF
            logging.basicConfig(level=logging.INFO)
            logging.info("File compressed successfully: %s", file_ID)
            pdf_compressed.save(compressed_pdf_path)

        fs.delete(ObjectId(file_ID))
        # Upload the compressed file back to MongoDB
        with open(compressed_pdf_path, "rb") as compressed_file:
            fs.upload_from_stream_with_id(
                ObjectId(file_ID), compressed_file.name, source=compressed_file
            )
        logging.basicConfig(level=logging.INFO)
        logging.info("Uploaded compressed file: %s", compressed_file.name)

    except Exception as e:
        logging.basicConfig(level=logging.ERROR)
        logging.error(f"Error during compression logic: {e}")
