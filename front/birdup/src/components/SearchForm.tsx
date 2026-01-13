import { FormEvent } from "react";
import { Bird, Search } from "lucide-react";

export interface SearchFormProps {
  searchTerm: string;
  setSearchTerm: (s: string) => void;
  loading: boolean;
  handleSearchSubmit: (e: FormEvent) => Promise<void>;
}

export default function SearchForm({
  searchTerm,
  setSearchTerm,
  loading,
  handleSearchSubmit,
}: SearchFormProps) {
  return (
    <form onSubmit={handleSearchSubmit}>
      <div className="flex items-center max-w-xl mx-auto space-x-2">
        <div className="relative w-full">
          <label htmlFor="searchInput">Region code</label>
          <div className="absolute flex items-center ps-3 pt-2 pointer-events-none">
            <Bird className="text-slate-500" strokeWidth={1}/>
          </div>
          <input
            type="text"
            id="searchInput"
            value={searchTerm}
            className="rounded-md py-2 px-10 w-full bg-transparent border-slate-400 border shadow-sm focus:ring-brand focus:border-slate-100 focus:shadow-md placeholder:text-slate-400" required
            placeholder={`Search by region code (e.g. US-LA)`}
            onChange={(e) => {
              // Set the internal search term to match the text input's contents
              setSearchTerm(e.target.value);
            }}
            disabled={loading}
          />
        </div>
        <button 
          className="bg-blue-300 rounded-md px-4 py-4 cursor-pointer" 
          type="submit"
          disabled={loading}
          >
          <Search className="text-white w-5 h-5"/>
        </button>
      </div>
    </form>
  );
}
