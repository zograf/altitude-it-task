import "./PopUp.css"
import { useEffect, useRef, useState } from "react";

export const usePopup = () => {
    const [isVisible, setIsVisible] = useState(false)
    const [args, setArgs] = useState(undefined)


    const showPopup = (args) => {
        setIsVisible(true)
        setArgs(args)
    }

    const onCloseCallback = () => {
        return () => {
            setIsVisible(false)
        }
    }

    return {
        isVisible,
        onCloseCallback,
        showPopup,
    }
}

export function PopUpFrame({visible, children}) {
    const dialog = useRef(null)

    useEffect(() => {
        if (dialog.current?.open && !visible) dialog.current?.close()
        else if (!dialog.current?.open && visible) dialog.current?.showModal()
    }, [visible])

    return(
        <dialog className="card" ref={dialog}>
            {children}
        </dialog>
    )
}

export function PopUpPage({visible, children}) {
    const dialog = useRef(null)

    useEffect(() => {
        if (dialog.current?.open && !visible) dialog.current?.close()
        else if (!dialog.current?.open && visible) dialog.current?.showModal()
    }, [visible])

    return(
        <dialog className="page-dialog w-100" ref={dialog}>
            {children}
        </dialog>
    )
}

export function StandardPopUp({visible, title, description, children, closeCallback}) {
    return(
        <PopUpFrame visible={visible}>
            <div style={{padding: "0 6px"}}>
                <p className="card-title v-spacer-xs" >{title}</p>
                <p className="card-body v-spacer-s">{description}</p>
            </div>
            {children}
            <button className="solid-button" onClick={closeCallback()}>Close</button>
        </PopUpFrame>
    )
}

export function MessagePopUp({popup, title, message}) {
    return(
        <StandardPopUp
            visible={popup.isVisible}
            title={title}
            description={message}
            closeCallback={popup.onCloseCallback}
        />
    )
}