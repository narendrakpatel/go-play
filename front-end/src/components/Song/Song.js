import React from 'react';

const Song = ({name, email}) => {
    return (
        <div className='tc dib br3 pa3 ma2 grow bw2 shadow-5'><label>{name}</label><br/>
        <label>{email}</label>
        </div>
    )
}

export default Song;