package main

import (
    "fmt"
    "strconv"
    "time"
    "github.com/jung-kurt/gofpdf"
    "strings"
)

func generador_factura() {
    var nombre_cliente, producto, cantidad_productos, precio_unitario, respuesta string
    fecha_factura := time.Now() 
    var total_general_factura float64

    // Crear una lista para almacenar los productos y sus detalles
    var productos [][]string

    fmt.Print("Ingresa tu nombre: ")
    fmt.Scanln(&nombre_cliente)

    for {
        fmt.Println("Ingresa el nombre del producto que vas a comprar: ")
        fmt.Scanln(&producto)
       

        fmt.Println("Ingresa la cantidad que deseas comprar: ")
        fmt.Scanln(&cantidad_productos)
       

        fmt.Println("Ingresa el precio unitario del producto: ")
        fmt.Scanln(&precio_unitario)
       

        // Conversión de strings a valores numéricos
        cantidad, err := strconv.Atoi(cantidad_productos) // La convierte a entero
        if err != nil {
            fmt.Println("Error: La cantidad ingresada no es válida.")
            return
        }

        precio, err := strconv.ParseFloat(precio_unitario, 64) // Convertir precio a float64
        if err != nil {
            fmt.Println("Error: El precio ingresado no es válido.")
            return
        }

        // Calcular el costo total del producto
        costo_total_producto := float64(cantidad) * precio
        total_general_factura += costo_total_producto

        // Almacenar los detalles del producto en el arreglo "productos []"
        productos = append(productos, []string{producto, cantidad_productos, fmt.Sprintf("%.2f", precio), fmt.Sprintf("%.2f", costo_total_producto)})

        // Preguntar si el cliente desea agregar más productos
        fmt.Println("¿Deseas agregar más productos? (sí/s/no): ")
        fmt.Scanln(&respuesta)
     

        respuesta = strings.ToLower(respuesta) // Convertir la respuesta a minúsculas en caso de que el cliente ingrese en mayúsculas para poder comparar
        fmt.Println("respuesta ingresada " + respuesta)
        if respuesta == "no" {
            break
        }
    }

    // Imprimir resumen de la factura por consola
    fmt.Println("\n *_:_* Resumen de la Factura *_:_*")
    fmt.Println("Cliente:", nombre_cliente)
    fmt.Println("Fecha de Factura:", fecha_factura.Format("2006-01-02"))
    fmt.Printf("Total General: %.2f\n", total_general_factura)


    // Sección para la conversión a PDF

    // Generar el archivo PDF con la factura
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)

    // Título de la factura
    pdf.Cell(40, 10, "Factura de compra tienda 'Jadeneitor Stores'")
    pdf.Ln(12)

    // Información del cliente y fecha
    pdf.SetFont("Arial", "", 12)
    pdf.Cell(40, 10, "Nombre del cliente: " + nombre_cliente)
    pdf.Ln(10)
    pdf.Cell(40, 10, "Fecha: " + fecha_factura.Format("2006-01-02"))
    pdf.Ln(10)

    // Crear la tabla de productos
    pdf.SetFont("Arial", "B", 12)
    pdf.CellFormat(40, 10, "Producto", "1", 0, "", false, 0, "")
    pdf.CellFormat(40, 10, "Cantidad", "1", 0, "", false, 0, "")
    pdf.CellFormat(40, 10, "Precio Unitario", "1", 0, "", false, 0, "")
    pdf.CellFormat(40, 10, "Total", "1", 1, "", false, 0, "") // El '1' al final hace que el cursor salte a la siguiente línea

    // Agregar los productos a la tabla
    pdf.SetFont("Arial", "", 12)
    for _, prod := range productos {
        pdf.CellFormat(40, 10, prod[0], "1", 0, "", false, 0, "")   // Producto
        pdf.CellFormat(40, 10, prod[1], "1", 0, "", false, 0, "")   // Cantidad
        pdf.CellFormat(40, 10, prod[2], "1", 0, "", false, 0, "")   // Precio Unitario
        pdf.CellFormat(40, 10, prod[3], "1", 1, "", false, 0, "")   // Total
    }

    // Total general
    pdf.Ln(10)
    pdf.CellFormat(40, 10, fmt.Sprintf("Total General: %.2f", total_general_factura), "", 1, "", false, 0, "")

    // Guardar el PDF
    err := pdf.OutputFileAndClose("factura.pdf")
    if err != nil {
        fmt.Println("Error al generar el PDF:", err)
    } else {
        fmt.Println("Factura generada exitosamente en factura.pdf")
    }
}

func main() {
    generador_factura()
}
