import { FreeTile, PlayerTile, PokeTile } from "./TileSVG"

export interface TileProps {
    offsetX?: number
    offsetY?: number,
    key?: string,
    children?: React.ReactNode
}

export const Tile = (props: TileProps) => {
    const {
        offsetX,
        offsetY,
        key,
        children
    } = props

    return (
        <>
            {   
                (children == "FT" && <FreeTile {...props}/>) ||
                (children == "&" && <PokeTile {...props}/>) ||
                (children == "#" && <PlayerTile {...props}/>)
            }
        </>
    )
}