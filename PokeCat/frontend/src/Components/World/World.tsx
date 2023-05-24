
import { Button, Col, Image, Input, Modal, Row, Typography } from "antd"

import {  useEffect, useRef, useState } from "react"
import { getWorld, removePlayer, sendAction } from "../../API"
import { Tile } from "../SVG Components"
import { useIdleTimer } from "react-idle-timer"

export interface PokeDex {
    [ID: string]: Pokemon
}

interface Pokemon {
    id:               number
	type:             string
	img_link:              string
	name:             string
	base_experience:   number
	effort_value_yield: number[]
	form:             string[]
	attack:           number
	defense:          number
	special_attack:    number
	special_defense:   number
	speed:            number
	max_hp:            number
}


export const World = (props: {setUsername: (username: string) => void, username: string}) => {
    
    const map = useRef<SVGGElement>(null)
    const [worldGrid, setWorldGrid] = useState<Array<Array<string>>>()
    const [playerCoord, setPlayerCoord] = useState<number[]>([])
    const [pokeDex, setPokeDex] = useState<PokeDex>()
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [pokeID, setPokeID] = useState<string>("")
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
            getWorld(props.username, setWorldGrid, setPlayerCoord, map.current!, setPokeDex)
        }, 500)

        const handleKeydown = (e: KeyboardEvent) => {
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
    const handleSearch = (value: string) => {
        setPokeID(value)
    }


    return (
        <>
            <Row justify={"center"}>
                <Col>
                    <Button onClick={() => setIsOpen(true)}>Open PokeDex</Button>
                    <Modal open={isOpen} onCancel={() => setIsOpen(false)} onOk={() => setIsOpen(false)}>
                        <Typography.Title level={2}>Pokedex</Typography.Title>
                        <Input.Search onSearch={(value) => handleSearch(value)}></Input.Search>

                        {pokeID && pokeDex && pokeDex[pokeID] && (
                            <>
                                <Image src={pokeDex[pokeID]["img_link"]}/>
                                <p>Name: {pokeDex[pokeID]["name"]}</p>
                                <p>ID: {pokeDex[pokeID]["id"]}</p>
                                <p>Type: {pokeDex[pokeID]["type"]}</p>
                                <p>Base XP: {pokeDex[pokeID]["base_experience"]}</p>
                                <p>Total EV: {pokeDex[pokeID]["effort_value_yield"]}</p>
                                <p>Form: {pokeDex[pokeID]["form"]}</p>
                                <p>Atk: {pokeDex[pokeID]["attack"]}</p>
                                <p>Def: {pokeDex[pokeID]["defense"]}</p>
                                <p>Special Atk: {pokeDex[pokeID]["special_attack"]}</p>
                                <p>Special Def: {pokeDex[pokeID]["special_defense"]}</p>
                                <p>Speed: {pokeDex[pokeID]["speed"]}</p>
                                <p>Max HP: {pokeDex[pokeID]["max_hp"]}</p>
                            </>
                        )}
                    </Modal>
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