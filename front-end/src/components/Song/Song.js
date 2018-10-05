import React from 'react';
import './Song.css'

const Song = ({ title, url, thumb, onSongClicked }) => {
    return (
        <div className='tc dib br3 pa3 ma2 grow bw2 shadow-5'
            onClick={() => onSongClicked({ url })}>
            <img src={thumb} alt={title} className='dimention' /><br />
            <label className='label'>{title}</label><br />
        </div>
    )
}

export default Song;