import type { birdObservation } from "../../types/shared.types";
import { MapPin, Calendar } from "lucide-react";

export interface BirdCardProps {
  observation: birdObservation;
}

export default function BirdCard({ observation }: BirdCardProps) {
  // Converts a raw date string from the backend to the format "mmm dd, yyyy"
  const fmtDate = (dateStr: string): string => {
    const date = new Date(dateStr);
    
    // Specify format for date string display
    const opts: Intl.DateTimeFormatOptions = {
        weekday: undefined, // Exclude day of the week
        year: "numeric", 
        month: "short", // e.g. "Jan", "Dec"
        day: "2-digit",
        hour: "2-digit",
        minute: "2-digit" 
    }
    
    return date.toLocaleDateString("en-US", opts);
  };

  // Consistent scale for svg icons
  const iconSize: number = 22;

  return (
    <div className="rounded-md p-6 flex flex-col gap-1 drop-shadow-md bg-white">
      <div className="flex justify-between">
        <h3 className="flex gap-2 text-lg">
          <b>{observation.comName}</b>  
          <span className="italic text-gray-600">{observation.sciName}</span>
        </h3>
        <span className="px-2 py-1 text-md self-start ml-1 bg-green-300 text-green-600 shadow-sm rounded">{observation.howMany}</span>
      </div>

      <div className="flex flex-col grow text-sm space-y-2 justify-end">
        <div className="flex gap-1.5"><MapPin size={iconSize} /><span className="truncate">{observation.locName}</span></div>
        <div className="flex gap-1.5"><Calendar size={iconSize} /> {fmtDate(observation.obsDt)} </div>
      </div>
    </div>
  );
}
