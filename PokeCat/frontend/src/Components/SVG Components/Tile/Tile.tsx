import { FreeTile, PlayerTile, PokeTile } from "./TileSVG"

export interface TileProps {
    offsetX?: number
    offsetY?: number,
    key?: string,
    children?: string
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
                (children?.includes("&") && <PokeTile {...props}/>) ||
                (children == "#" && <PlayerTile {...props}/>)
            }
        </>
    )
}