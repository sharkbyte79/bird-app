import { birdObservation } from "../../types/shared.types";

export interface BirdCardProps {
  observation: birdObservation;
}

export default function BirdCard({observation} :BirdCardProps) {
    return (
        <div className="rounded-s p-4 flex">
            {observation.comName}
        </div>
    );
}
