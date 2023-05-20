import { Col, Row } from "antd"
import {  useEffect, useRef, useState } from "react"
import { getWorld, sendAction } from "../../API"
import { Tile } from "../SVG Components"

export const World = () => {
    
    const map = useRef<SVGGElement>(null)
    
    useEffect(() => {
        getWorld(setWorldGrid)

        const handleKeydown = (e: KeyboardEvent) => {
            e.preventDefault()
            sendAction(e.key, setKeydown)

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

        document.addEventListener('keydown', handleKeydown)

        return () => {
            document.removeEventListener('keydown', handleKeydown)
        }
    }, [])

    const [keydown, setKeydown] = useState("")
    const [worldGrid, setWorldGrid] = useState<Array<Array<string>>>()

    return (
        <>
            <Row justify={"center"}>
                <Col>
                    
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