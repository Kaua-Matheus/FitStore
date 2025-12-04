import { useState } from "react"

import Button from "./Button";

export default function Carousel() {

    const images = [
        "https://picsum.photos/id/1/800/400",
        "https://picsum.photos/id/2/800/400",
        "https://picsum.photos/id/3/800/400"
    ]

    var [index, setIndex] = useState(0);

    const next = () => {setIndex((index + 1) % images.length)};
    const prev = () => {setIndex((index - 1 + images.length) % images.length)};

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