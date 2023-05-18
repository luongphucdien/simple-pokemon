import { TileProps } from "../Tile"

export const FreeTile = (props: TileProps) => {
    const {
        offsetX,offsetY
    } = props
    return(
        <g>
            <rect x={offsetX} y={offsetY} width="40" height="40" fill="blue" stroke="white"/>
            {/* <line x1={offsetX} x2={offsetX!+40} y1={offsetY} y2={offsetY!+40} stroke="red"/>
            <line x1={offsetX!+40} x2={offsetX} y1={offsetY} y2={offsetY!+40} stroke="red"/> */}
        </g>
    )
}