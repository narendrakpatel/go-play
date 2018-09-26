import React from 'react';
import Song from '../Song/Song';

const SongQueue = ({searchResults}) => {
    return (
        searchResults.map((item, id) => {
            return <Song key={id} name={item.name} email={item.email}/>
        })
    )
}

export default SongQueue;