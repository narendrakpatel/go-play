import React from 'react';
import Song from '../Song/Song';

const SongQueue = ({searchResults, onSongClicked}) => {
    return (
        searchResults.map((item, id) => {
            return <Song key={id} title={item.title} url={item.url} thumb={item.thumb} onSongClicked={onSongClicked} />
        })
    )
}

export default SongQueue;