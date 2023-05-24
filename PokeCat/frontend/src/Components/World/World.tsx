import { Col, Row } from "antd"
import {  useEffect, useRef, useState } from "react"
import { getWorld, removePlayer, sendAction } from "../../API"
import { Tile } from "../SVG Components"
import { useIdleTimer } from "react-idle-timer"

export const World = (props: {setUsername: (username: string) => void, username: string}) => {
    
    const map = useRef<SVGGElement>(null)
    const [worldGrid, setWorldGrid] = useState<Array<Array<string>>>()
    const [playerCoord, setPlayerCoord] = useState<number[]>([])

    const onIdle = () => {
        removePlayer(props.username)
        props.setUsername("")
      }
    
      useIdleTimer({
        onIdle,
        timeout: 1800_000,
      })
    
    useEffect(() => {
        // let t: SVGTransform
        // let g: SVGGElement = map.current!
        
        const interval = setInterval(() => {
            getWorld(props.username, setWorldGrid, setPlayerCoord, map.current!)
        }, 500)

        const handleKeydown = (e: KeyboardEvent) => {
            e.preventDefault()
            sendAction(e.key, props.username)
        }

        const handleReload = () => {
            removePlayer(props.username)
            props.setUsername("")
        }

        document.addEventListener('keydown', handleKeydown)
        window.addEventListener('beforeunload', handleReload)

        return () => {
            document.removeEventListener('keydown', handleKeydown)
            window.removeEventListener('beforeunload', handleReload)
            clearInterval(interval)
        }
    }, [])

    

    return (
        <>
            <Row justify={"center"}>
                <Col>
                    <p>{props.username}</p>
                    <p>{`[${playerCoord[0]}, ${playerCoord[1]}]`}</p>
                    
                    <svg viewBox="0 0 400 400" width={400} height={400}>
                        <line x1={200} x2={200} y1={0} y2={400} stroke="red"></line>
                        <line x1={0} x2={400} y1={200} y2={200} stroke="red"></line>

                        <rect width="100%" height="100%" x="0" y="0" fill="gray" opacity={0.5}/>
                        
                        <g ref={map}>
                            {worldGrid?.map((y,i)=>(
                                y.map((x,j)=>(
                                    <Tile offsetX={j*40} offsetY={i*40}>{x}</Tile>
                                ))
                            ))}
                        </g>
                        
                    </svg>
                </Col>
            </Row>
        </>
    )
}