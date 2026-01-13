import type { birdObservation } from "../../types/shared.types";
import BirdCard from "./BirdCard";

export interface SearchResultsProps {
  observations: birdObservation[];
  regionCode: string;
}

export default function SearchResults({ observations, regionCode }: SearchResultsProps) {
  return (
    <>
      <h2 className="text-xl font-bold">Found {observations.length} {observations.length === 0 ? "bird" : "birds"}</h2>
      <p className="text-slate-600">Region: {regionCode.toUpperCase()}</p>
      {/* Arrange cards with 1, 2 or 3 column display relative to page size */}
      <div className="grid sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 auto-rows-fr mt-3">
        {observations.length > 0 &&
          // Map every observation to a bird card
          observations.map((obs, idx) => <BirdCard key={idx} observation={obs} />)}
      </div>
    </>

  );
}
