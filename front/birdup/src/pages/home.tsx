import React, { useEffect, useState } from "react";
import { birdObservation } from "../../types/shared.types.ts";
import { birdupResponse, getObservations } from "../services/apiClient.ts";
import SearchForm from "../components/SearchForm.tsx";
import LoadingSpinner from "../components/LoadingSpinner.tsx";
import SearchResults from "../components/SearchResults.tsx";
import ErrorDisplay from "../components/ErrorDisplay.tsx";

export default function Home() {
  const [observations, setObservations] = useState<birdObservation[]>([]);
  const [notable, setNotable] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  const [firstLoad, setFirstLoad] = useState<boolean>(true);
  const [error, setError] = useState<string>("");
  const [searchTerm, setSearchTerm] = useState<string>("");

  async function handleSearchSubmit(e: React.FormEvent): Promise<void> {
    // Prevent page reload on search form submission
    e.preventDefault();

    try {
      setFirstLoad(false);
      setLoading(true);
      const res = await getObservations(searchTerm.trim().toUpperCase(), notable);
      setObservations(res.data);
    } catch (err) {
      const errMsg: string = `${err}`;
      setError(errMsg);
    } finally {
      setLoading(false);
    }
  }

  return (
    <>
      <SearchForm
        searchTerm={searchTerm}
        setSearchTerm={setSearchTerm}
        loading={loading}
        handleSearchSubmit={handleSearchSubmit}
      />

      {error.length > 0 && <ErrorDisplay text={error} />}

      {loading && <LoadingSpinner message="Fetching " />}

      {observations.length > 0 && 
      <>      
        <SearchResults observations={observations} regionCode={searchTerm}/>
      </>
      }
    </>
  );
}
