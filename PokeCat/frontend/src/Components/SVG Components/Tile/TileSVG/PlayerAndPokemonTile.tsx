import { TileProps } from "../Tile"

export const PlayerAndPokemonTile = (props: TileProps) => {
    const {
        offsetX,offsetY
    } = props
    return(
        <g>
            <rect x={offsetX} y={offsetY} width="40" height="40" fill="red" stroke="white"/>
            <circle r="10" fill="green" cx="20" cy="20"/>
            {/* <line x1={offsetX} x2={offsetX!+40} y1={offsetY} y2={offsetY!+40} stroke="red"/>
            <line x1={offsetX!+40} x2={offsetX} y1={offsetY} y2={offsetY!+40} stroke="red"/> */}
        </g>
    )
}