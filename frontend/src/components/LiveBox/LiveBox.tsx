import { LiveBoxContainer } from "./LiveBoxContainer"
import { LiveBoxFooter } from "./LiveBoxFooter"
import { LiveBoxHeader } from "./LiveBoxHeader"






export const LiveBox = () => {
    return (
        <>
            <LiveBoxHeader />
            <LiveBoxContainer></LiveBoxContainer>
            <LiveBoxFooter />
        </>
    )
}