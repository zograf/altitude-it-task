import {useEffect, useState} from "react";
import "./DropDown.css"

export function DropDownInput({options, icon, placeholder, callback}) {
    const [_, setSelected] = useState(undefined)
    const [value, setValue] = useState("")
    const [isVisible, setIsVisible] = useState(false)
    const [showError, setError] = useState(false)

    useEffect(() => {
        if(options.length === 0) {
            setError(value.length !== 0)
            setValue("")
            setSelected(undefined)
        }
    }, [options.length])

    const handleFocus = (e) => {
        setTimeout(() => {
            setIsVisible(e.type === 'focus')
            if (isVisible) setError(options.find((item) => item.label === e.target.value) === undefined)
        }, 200)
    }
    const itemSelected = (item) => {
        setError(false)
        setSelected(item)
        setValue(item.label)
        callback(item.value)
    }
    const handleValueChanged = (e) => {
        let possibleItem = options.find((item) => item.label === e.target.value)
        if(possibleItem === undefined) {
            setSelected(undefined)
            callback(undefined)
        }
        else {
            setSelected(possibleItem)
            callback(possibleItem.value)
        }
        setValue(e.target.value)
    }

    return(
        <div className={['input-wrapper v-spacer-xs', showError ? 'error-border' : "regular-border", options.length === 0 ? 'disabled-input' : ''].join(' ')}>
            <span className="material-symbols-outlined icon input-icon">{icon}</span>
            <input className={['select-input'].join(' ')} disabled={options.length === 0} placeholder={placeholder} value={value} onChange={handleValueChanged} onFocus={handleFocus} onBlur={handleFocus}/>
            <div className={['drop-down', isVisible ? 'visible' : ""].join(' ')}>
                { options?.filter((item) => {return item.label.toLowerCase().includes(value.toLowerCase())}).map((item) => {
                    return( <li onClick={(e) => { itemSelected(item) }}>{item.label}</li> )
                })}
            </div>
        </div>
    ) 
}

export function DropDownSelect({options, icon, placeholder, callback}) {
    const [value, setValue] = useState("")
    const [isVisible, setIsVisible] = useState(false)
    const [showError, setError] = useState(false)

    useEffect(() => {
        if(options.length === 0) {
            setError(value.length !== 0)
            setValue("")
        }
    }, [options.length])

    const handleOnClick = () => {
        //console.log("Click")
        setIsVisible(!isVisible)
    }
    const hide = () => {
        setIsVisible(false)
    }
    const itemSelected = (item) => {
        setError(false)
        setValue(item.label)
        callback(item.value)
        setIsVisible(false)
    }
    const handleSelect = (e) => {
        console.log(e)
    }

    return(
        <div className={['input-wrapper', showError ? 'error-border' : "regular-border", options.length === 0 ? 'disabled-input' : ''].join(' ')} onMouseLeave={hide} style={{maxWidth: "150px"}}>
            <span className="material-symbols-outlined icon input-icon">{icon}</span>
            <input className={['select-input'].join(' ')} style={{userSelect:"none"}} contentEditable={false} readOnly={true} placeholder={placeholder} value={value} onClick={handleOnClick} onSelect={handleSelect}/>
            <div className={['drop-down', isVisible ? 'visible' : ""].join(' ')}>
                { options?.map((item) => {
                    return( <li onClick={(e) => { itemSelected(item) }}>{item.label}</li> )
                })}
            </div>
        </div>
    )
}