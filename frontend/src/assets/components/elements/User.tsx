// React
import { useState, useEffect, useRef, type RefObject, type ReactNode } from "react"

// React Icons
import { CiUser } from "react-icons/ci";
import { IoKeyOutline, IoEyeOutline, IoEyeOffOutline, IoCloseCircleOutline } from "react-icons/io5";
import { FaGoogle } from "react-icons/fa";

// Componentes
import Button from "./Button";
import { useToast } from '../../context/useToast'
import { useAuth } from "../../context/useAuth";


interface User {
    login: string;
    password: string;
}

// Definindo o Enum para valores fixos
const AccessMode = {
    LOGIN: "login",
    LOGON: "logon",
    LOGOUT: "logout",
} as const;

type AccessMode = typeof AccessMode[keyof typeof AccessMode];


function UserBox({ children }: { children?: ReactNode }) { // Melhorar funcionamento (ainda é necessário aprimorar esquema de fechar)
    return (
        <div
            className="fixed top-0 right-0 bg-custom-gray bg-opacity-50 
            w-[100vw] h-[100vh] transform">
            <div className="flex justify-center mt-[15vh]">
                <div className="flex flex-col items-center justify-center bg-custom-light-gray w-[40vw] h-[50vh] space-y-6 rounded-sm shadow-md">

                    { children } 

                </div>
            </div>
        </div>
    )
}

export default function User() {

    const { addToast } = useToast();
    const { checkAuth, isLoggedIn, user } = useAuth(); // user para adquirir informações do usuário

    const [opened, setOpened] = useState(false)
    const [hidden, setHidden] = useState(true)

    const [access, setAccess] = useState<AccessMode>(AccessMode.LOGIN)

    // User Credentials
    const inputUser = useRef<HTMLInputElement>(null)
    const inputLogin = useRef<HTMLInputElement>(null)
    const inputPassword = useRef<HTMLInputElement>(null)

    // Image
    const [profileImage, setProfileImage] = useState<string | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(true)

    useEffect(() => {
        if (opened) {
            document.body.style.overflow = "hidden";
        } else {
            document.body.style.overflow = "";
        }

        const fetchAll = async () => {
            const image = await getImage();
            setProfileImage(image)
        }

        if (isLoggedIn && profileImage == undefined) {
            fetchAll()
        }
    })

    const userRegister = async (login: RefObject<HTMLInputElement | null>, password: RefObject<HTMLInputElement | null>) => {
        try {
            var parsedLogin = login.current != null ? login.current.value : ""
            var parsedPassword = password.current != null ? password.current.value : ""

            await fetch("http://localhost:8080/user/register", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify({ "user_name": parsedLogin, "login": parsedLogin, "password": parsedPassword })
            })
            setOpened(false);
            addToast("Registro executado com sucesso, bem-vindo!", "success");
            checkAuth();
            console.log(isLoggedIn)
        } catch (err) {
            console.log(`Error: ${err}`)
        }
    }

    const loginUser = async (login: RefObject<HTMLInputElement | null>, password: RefObject<HTMLInputElement | null>) => {
        try {
            var parsedLogin = login.current != null ? login.current.value : ""
            var parsedPassword = password.current != null ? password.current.value : ""

            const response = await fetch("http://localhost:8080/user/login", {
                method: "POST",
                headers: { "Content-Type": "application/json", },
                credentials: "include",
                body: JSON.stringify({ "user_name": parsedLogin, "login": parsedLogin, "password": parsedPassword })
            })

            const data = await response.json();

            if (data.authenticated) {
                setOpened(false);
                addToast("Login executado com sucesso", "success");
                checkAuth();
                console.log(isLoggedIn)
            } else {
                addToast("Senha incorreta", "error");
            }

            return;
        } catch (err) {
            console.log(`Error: ${err}`)
        }
    }

    const logoutUser = async () => {
        try {
            await fetch("http://localhost:8080/user/logout", {
                method: "GET",
                credentials: "include"
            })
            setOpened(false);
            addToast("Logout executado com sucesso", "success");
            checkAuth()
            console.log(isLoggedIn)
        } catch (err) {
            console.log(`Error: ${err}`)
        }
    }

    async function getImage() {
        try {
            setIsLoading(true)
            const response = await fetch(`http://localhost:8080/user/get/${user?.user_login}`, {
                method: "GET",
                credentials: "include",
            })

            const jsonResponse = await response.json()

            const image = await fetch(`http://localhost:8080/image/${jsonResponse.profile_image}`, {
                method: "GET",
                credentials: "include",
            })

            const jsonImage = await image.json()

            return `http://localhost:8080/files/${jsonImage.data.file_path}/${jsonImage.data.file_name}${jsonImage.data.content_type}` // Retirar do endpoint o data, mandando só o arquivo final
        } catch (err) {
            console.log(`Error: ${err}`)
        }
        finally {
            setIsLoading(false)
        }
    }

    function Icon() {
        if (isLoggedIn && !isLoading) {
            return (
                <img 
                width="40"
                height="40" // Arrumar tamanho da imagem
                className="hover:shadow-md cursor-pointer rounded-full" 
                src={profileImage} alt={user?.user_login} 
                onClick={() => { setOpened(!opened); setAccess(AccessMode.LOGOUT) }}/>
            )
        }
        return (
            <CiUser
                size={40} 
                className="
            text-custom-dark
            bg-custom-light-gray
            active:bg-custom-secondary transition-colors duration-400
            hover:shadow-md cursor-pointer
            text-center p-2 rounded"
                onClick={() => { setOpened(!opened); isLoggedIn ? setAccess(AccessMode.LOGOUT) : setAccess(AccessMode.LOGIN) }}
            />
        )
    }

    function Card(type?: "login" | "logon" | "logout") {
        switch (type) {
            case "login":
                return (
                    <UserBox>
                        <IoCloseCircleOutline
                            size={35}
                            onClick={() => { setOpened(!opened) }}
                            className="fixed top-[15.5vh] right-[30.5vw] cursor-pointer" /> { /*Adicionar valor variável em size*/}

                        <div className="flex flex-col space-y-2">
                            <h1 className="text-center font-bold">Acessar Conta</h1>
                            <div className="relative">
                                <CiUser className="absolute left-3 top-1/2 transform -translate-y-1/2 text-custom-dark" />
                                <input ref={inputLogin} className="px-10 w-[30vw] h-[30px] text-custom-dark border border-custom-gray rounded-sm" placeholder="Login" name="login" type="text" />
                            </div>
                            <div className="relative">
                                <IoKeyOutline className="absolute left-3 top-1/2 transform -translate-y-1/2 text-custom-dark" />
                                <input ref={inputPassword} className="px-10 w-[30vw] h-[30px] text-custom-dark border border-custom-gray rounded-sm" placeholder="Senha" name="password" type={hidden ? "password" : "text"} />
                                {hidden && (
                                    <IoEyeOutline
                                        className="absolute right-3 top-1/2 transform -translate-y-1/2 text-custom-dark cursor-pointer"
                                        onClick={() => { setHidden(!hidden) }}
                                    />
                                ) || (
                                        <IoEyeOffOutline
                                            className="absolute right-3 top-1/2 transform -translate-y-1/2 text-custom-dark cursor-pointer"
                                            onClick={() => { setHidden(!hidden) }}
                                        />)}

                            </div>
                            <p className="text-right text-sm">Esqueceu a senha? <span className="text-custom-primary underline cursor-pointer">Recuperar</span></p>
                        </div>
                        <div className="flex flex-col space-y-2">
                            <Button onClick={() => loginUser(inputLogin, inputPassword)} type="dark" className="w-[30vw]">Acessar</Button>
                            <Button onClick={() => setAccess(AccessMode.LOGON)} type="light" className="w-[30vw] font-light">Não possui conta? Criar agora</Button>
                        </div>
                        <div className="flex flex-col">
                            <p className="text-center">Ou</p>
                            <div className="relative">
                                <FaGoogle size={24} className="absolute left-3 top-1/2 transform -translate-y-1/2 text-custom-primary cursor-pointer" />
                                <button className={`
                                    w-[30vw]
                                    text-custom-dark bg-custom-light-gray font-light
                                    text-center p-2 rounded-sm shadow-md
                                    border border-custom-gray
                                    active:shadow-sm active:bg-custom-secondary
                                    transition-colors duration-400
                                    `}>
                                    Acesse com o Google
                                </button>
                            </div>
                        </div>
                    </UserBox>
                )

            case "logon":
                return (
                    <UserBox>
                        <IoCloseCircleOutline
                            size={35}
                            onClick={() => { setOpened(!opened) }}
                            className="fixed top-[15.5vh] right-[30.5vw] cursor-pointer" /> { /*Adicionar valor variável em size*/}

                        <div className="flex flex-col space-y-2">
                            <h1 className="text-center font-bold">Criar Conta</h1>
                            <div className="relative">
                                <CiUser className="absolute left-3 top-1/2 transform -translate-y-1/2 text-custom-dark" />
                                <input ref={inputUser} className="px-10 w-[30vw] h-[30px] text-custom-dark border border-custom-gray rounded-sm" placeholder="Nome do Usuário" name="login" type="text" />
                            </div>
                            <div className="relative">
                                <CiUser className="absolute left-3 top-1/2 transform -translate-y-1/2 text-custom-dark" />
                                <input ref={inputLogin} className="px-10 w-[30vw] h-[30px] text-custom-dark border border-custom-gray rounded-sm" placeholder="Login" name="login" type="text" />
                            </div>
                            <div className="relative">
                                <IoKeyOutline className="absolute left-3 top-1/2 transform -translate-y-1/2 text-custom-dark" />
                                <input ref={inputPassword} className="px-10 w-[30vw] h-[30px] text-custom-dark border border-custom-gray rounded-sm" placeholder="Senha" name="password" type={hidden ? "password" : "text"} />

                            </div>
                            <div className="relative">
                                <IoKeyOutline className="absolute left-3 top-1/2 transform -translate-y-1/2 text-custom-dark" />
                                <input ref={inputPassword} className="px-10 w-[30vw] h-[30px] text-custom-dark border border-custom-gray rounded-sm" placeholder="Repita a Senha" name="password" type={hidden ? "password" : "text"} />
                                {hidden && (
                                    <IoEyeOutline
                                        className="absolute right-3 top-1/2 transform -translate-y-1/2 text-custom-dark cursor-pointer"
                                        onClick={() => { setHidden(!hidden) }}
                                    />
                                ) || (
                                        <IoEyeOffOutline
                                            className="absolute right-3 top-1/2 transform -translate-y-1/2 text-custom-dark cursor-pointer"
                                            onClick={() => { setHidden(!hidden) }}
                                        />)}

                            </div>
                        </div>
                        <div className="flex flex-col space-y-2">
                            <Button onClick={() => userRegister(inputLogin, inputPassword)} type="dark" className="w-[30vw]">Criar</Button>
                            <Button onClick={() => setAccess(AccessMode.LOGIN)} type="light" className="w-[30vw] font-light">Possui conta? Acesse aqui</Button>
                        </div>
                        <div className="flex flex-col">
                            <p className="text-center">Ou</p>
                            <div className="relative">
                                <FaGoogle size={24} className="absolute left-3 top-1/2 transform -translate-y-1/2 text-custom-primary cursor-pointer" />
                                <button className={`
                                    w-[30vw]
                                    text-custom-dark bg-custom-light-gray font-light
                                    text-center p-2 rounded-sm shadow-md
                                    border border-custom-gray
                                    active:shadow-sm active:bg-custom-secondary
                                    transition-colors duration-400
                                    `}>
                                    Crie com o Google
                                </button>
                            </div>
                        </div>
                    </UserBox>
                )

            case "logout":
                return (
                    <UserBox>
                        <IoCloseCircleOutline
                            size={35}
                            onClick={() => { setOpened(!opened) }}
                            className="fixed top-[15.5vh] right-[30.5vw] cursor-pointer" /> { /*Adicionar valor variável em size*/}

                        <div className="flex flex-col space-y-2">
                            <h1 className="text-center font-bold">Mudar Conta / Sair</h1>
                        </div>

                        <div className="flex flex-col space-y-2">
                            <Button onClick={() => logoutUser()} type="dark" className="w-[30vw]">Sair</Button>
                            <Button onClick={() => setAccess(AccessMode.LOGIN)} type="light" className="w-[30vw] font-light">Mudar de Conta</Button>
                        </div>
                        <div className="flex flex-col">
                            <p className="text-center">Ou</p>
                            <div className="relative">
                                <FaGoogle size={24} className="absolute left-3 top-1/2 transform -translate-y-1/2 text-custom-primary cursor-pointer" />
                                <button className={`
                                    w-[30vw]
                                    text-custom-dark bg-custom-light-gray font-light
                                    text-center p-2 rounded-sm shadow-md
                                    border border-custom-gray
                                    active:shadow-sm active:bg-custom-secondary
                                    transition-colors duration-400
                                    `}>
                                    Acesse com o Google
                                </button>
                            </div>
                        </div>
                    </UserBox>
                )
        }
    }

    return (
        <>
            <Icon />

            {/* Componente de login */}
            {opened && (
                Card(access)
            )}
        </>
    )
}