interface LogoProps {
    darkmode?: boolean;
}

export default function Logo({darkmode=false}: LogoProps) {

    if (darkmode) {
        return (
            <div className="flex h-10 w-22 bg-custom-light-gray space-x-2 items-center rounded-sm px-2 shadow-md">
                <img className="h-12 rounded shadow-sm" src="src/assets/images/logos/Light.png" alt="Logo Light" />
                <p className="text-lg text-custom-dark"><span className="font-fugaz text-custom-primary">Fit</span>Store</p>
            </div>
        )
    } else {
        return (
            <div className="flex h-10 w-22 bg-custom-dark space-x-2 items-center rounded-sm px-2 shadow-md">
                <img className="h-12 shadow-sm rounded" src="src/assets/images/logos/Dark.png" alt="Logo Dark" />
                <p className="text-lg text-custom-light-gray"><span className="font-fugaz text-custom-primary">Fit</span>Store</p>
            </div>
        )
    }
}