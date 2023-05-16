import { Card } from "antd"
import { useEffect, useState } from "react"
import { getWorld, sendAction, testNoParams } from "../../API"

export const World = () => {
    useEffect(() => {
        getWorld()

        const handleKeydown = (e: KeyboardEvent) => {
            e.preventDefault()
            sendAction(e.key, setKeydown)
        }

        document.addEventListener('keydown', handleKeydown)

        return () => {
            document.removeEventListener('keydown', handleKeydown)
        }
    }, [])

    const [keydown, setKeydown] = useState("")

    return (
        <>
            <Card title="Key pressed" >
                <p>{keydown}</p>
            </Card>
        </>
    )
}