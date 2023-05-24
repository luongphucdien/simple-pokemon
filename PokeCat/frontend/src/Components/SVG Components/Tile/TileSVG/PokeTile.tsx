import { TileProps } from "../Tile"

export const PokeTile = (props: TileProps) => {
    const {
        offsetX,offsetY
    } = props
    return (
        <g>
            <rect x={offsetX} y={offsetY} width={40} height={40} fill="red"/>
            {/* <line x1="0" x2="40" y1="0" y2="40" stroke="black"/>
            <line x1="40" x2="0" y1="0" y2="40" stroke="black"/> */}
        </g>
    )
}