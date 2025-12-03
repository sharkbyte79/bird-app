import { FormEvent } from "react";

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
      <div className="flex flex-col gap-8">
          <label htmlFor="searchInput">Region code</label>
          <input
            type="text"
            id="searchInput"
            value={searchTerm}
            className="rounded-md p-3"
            placeholder={`Search by region code (e.g. US-LA)`}
            onChange={(e) => {
              setSearchTerm(e.target.value);
            }}
            disabled={loading}
          />
        <button className="bg-blue-400 rounded-md px-8 py-2 cursor-pointer" type="submit">
          Search
        </button>
      </div>
    </form>
  );
}
