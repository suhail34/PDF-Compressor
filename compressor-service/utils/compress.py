import fitz  # PyMuPDF
import shutil


def compress_pdf(input_path, output_path):
    # Open the PDF file
    pdf_document = fitz.open(input_path)
    # Create a new PDF with compressed content
    pdf_compressed = fitz.open()

    # Iterate through pages and add them to the new PDF
    for page_number in range(pdf_document.page_count):
        page = pdf_document[page_number]
        pix = page.get_pixmap()
        pdf_compressed.insert_page(
                page_number,
                width=pix.width,
                height=pix.height)
        pdf_compressed[page_number].insert_text((0, 0), page.get_text())

    # Save the compressed PDF
    pdf_compressed.save(output_path)
    pdf_document.close()
    pdf_compressed.close()
