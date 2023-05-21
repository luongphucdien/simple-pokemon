import { Col, Row } from "antd"
import {  useEffect, useRef, useState } from "react"
import { getWorld, removePlayer, sendAction } from "../../API"
import { Tile } from "../SVG Components"
import { useIdleTimer } from "react-idle-timer"

export const World = (props: {setUsername: (username: string) => void, username: string}) => {
    
    const map = useRef<SVGGElement>(null)

    const onIdle = () => {
        removePlayer(props.username)
        props.setUsername("")
      }
    
      useIdleTimer({
        onIdle,
        timeout: 1800_000,
      })
    
    useEffect(() => {
        
        const interval = setInterval(() => {
            getWorld(setWorldGrid)
        }, 500)

        const handleKeydown = (e: KeyboardEvent) => {
            e.preventDefault()
            sendAction(e.key, props.username)

            let t: SVGTransform
            let g:SVGGElement = map.current!

            if (e.key.toLowerCase() === "w") {
                if(g.transform.baseVal.numberOfItems === 0) {
                    g.setAttribute("transform", "translate(" + 0 + "," + (40) + ")")
                }
                else {
                    t = g.transform.baseVal.getItem(0)
                    t.setMatrix(t.matrix.translate(0, 40))
                }
            }
            else if (e.key.toLowerCase() === "s") {
                if(g.transform.baseVal.numberOfItems === 0) {
                    g.setAttribute("transform", "translate(" + 0 + "," + (-40) + ")")
                }
                else {
                    t = g.transform.baseVal.getItem(0)
                    t.setMatrix(t.matrix.translate(0, -40))
                }
            }
            else if (e.key.toLowerCase() === "a") {
                if(g.transform.baseVal.numberOfItems === 0) {
                    g.setAttribute("transform", "translate(" + (40) + "," + 0 + ")")
                }
                else {
                    t = g.transform.baseVal.getItem(0)
                    t.setMatrix(t.matrix.translate(40, 0))
                }
            }
            else if (e.key.toLowerCase() === "d") {
                if(g.transform.baseVal.numberOfItems === 0) {
                    g.setAttribute("transform", "translate(" + (-40) + "," + 0 + ")")
                }
                else {
                    t = g.transform.baseVal.getItem(0)
                    t.setMatrix(t.matrix.translate(-40, 0))
                }
            }
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

    const [worldGrid, setWorldGrid] = useState<Array<Array<string>>>()

    return (
        <>
            <Row justify={"center"}>
                <Col>
                    <p>{props.username}</p>
                    
                    <svg viewBox="0 0 400 400" width={400} height={400}>
                        <rect width="100%" height="100%" x="0" y="0" fill="gray"/>
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