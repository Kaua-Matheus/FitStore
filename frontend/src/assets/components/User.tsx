// React
import { useState, useEffect } from "react"

// React Icons
import { CiUser } from "react-icons/ci";

export default function User() {

    const [Opened, setOpened] = useState(Boolean)

    useEffect(() => {
        if (Opened) {
            document.body.style.overflow = "hidden";
        } else {
            document.body.style.overflow = "";
        }
    })

    return (
        <>
        <CiUser 
        size={40} className="
            text-custom-dark
            bg-custom-light-gray
            active:bg-custom-secondary transition-colors duration-400
            hover:shadow-md cursor-pointer
            text-center p-2 rounded"
        onClick={() => {setOpened(!Opened)}}
            />
            {Opened && (
                <div 
                className="fixed top-[40vh] left-[48.42vw]">
                    SOS
                </div>
            )}
        </>
    )
}