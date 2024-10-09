import { useEffect, useRef } from 'react'

export const BackgroundBanner  = ()=>{
    const ref =  useRef(null)

    useEffect(() => {
        const body = ref.current
        let compensate = window.innerWidth > 1650 ? window.innerWidth / 1650 : 1;
        let layout = []


    }, [])

    return (
        <>
            <div className={"background_banner"} ref={ref}>
            </div>
        </>
    )
}