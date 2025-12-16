import { useEffect, useState } from "react"

import Button from "./Button";

type Image = {
    filename: string
    url: string
}

export default function Carousel() {

    var [images, setImages] = useState<Image[]>([])
    var [index, setIndex] = useState(0);

    const next = () => {setIndex((index + 1) % images.length)};
    const prev = () => {setIndex((index - 1 + images.length) % images.length)};

    useEffect(() => {

        const fetchAll = async () => {

            try {
            const response = await fetch("http://localhost:8080/Banners");
            const data_image = await response.json();

            setImages(data_image.files);

            } catch(err) {
                console.log(`Erro: ${err}`);
            }
        };

        fetchAll();

    }, [])

    return (
        <div className="flex items-center space-x-4">
            <Button onClick={prev}>
                Prev
            </Button>

            {/* Adicionar descrição de imagens, podemos colocar isso em uma requisição http */}
            <img src={`${images[index]}`} alt="imagem" />

            <Button onClick={next}>
                Prox
            </Button>
        </div>
    )
}