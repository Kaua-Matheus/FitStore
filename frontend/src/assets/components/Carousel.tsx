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


            setImages(data_image.data);

            } catch(err) {
                console.log(`Erro: ${err}`);
            }
        };

        fetchAll();

    }, [])

    if (images.length === 0) {
        return (
            <div>
                Carregando...
            </div>
        )
    }

    return (
        <div className="flex items-center space-x-4">
            <Button onClick={prev}>
                Prev
            </Button>

            <div className="flex 
                md:h-[24vh] md:h-[40vh]">
                <img className="rounded-sm md:rounded-md lg:rounded-lg" src={`${images[index].url == undefined ? "" : images[index].url}`} alt={`${images[index].filename}`} />
            </div>

            <Button onClick={next}>
                Prox
            </Button>
        </div>
    )
}