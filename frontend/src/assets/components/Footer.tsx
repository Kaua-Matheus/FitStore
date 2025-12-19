export default function Footer() {
    return (
        <footer className="flex flex-col items-center border-2 space-y-[2vh]">
            <div className="flex space-x-[10vw]">
                <div>
                    <h1>Institucional</h1>
                    <p>Nosso Manifesto</p>
                    <p>Trabalhe Conosco</p>
                </div>
                <div>
                    <h1>Categorias</h1>
                    <p>Nosso Manifesto</p>
                    <p>Trabalhe Conosco</p>
                </div>
                <div className="space-y-8">
                    <div>
                        <h1>Atendimentos</h1>
                        <p>Nosso Manifesto</p>
                        <p>Trabalhe Conosco</p>
                    </div>
                    <div>
                        <h1>Redes</h1>
                        <p>Nosso Manifesto</p>
                        <p>Trabalhe Conosco</p>
                    </div>
                </div>
            </div>

            <div>
                Logo
            </div>

            <div className="font-outfit">
                <p>Tecnologia <span className="font-bold underline decoration-custom-primary">Kaua Matheus</span></p>
            </div>
        </footer>
    )
}