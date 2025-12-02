export default function Button() {
    return (
        <button className="bg-custom-secondary p-1 rounded" 
        onClick={() => {
            console.log("Botão clicado!");
        }}>
            Button!
        </button>
    )
}