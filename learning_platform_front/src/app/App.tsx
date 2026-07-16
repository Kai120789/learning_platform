import { useEffect, type FC } from 'react'
import { AppRouter } from './router/AppRouter'
import { useAppDispatch } from './providers/storeProvider/hooks/hooks'
import { getUserData } from '@/entities/user'

const App: FC = () => {
    const dispatch = useAppDispatch()

    useEffect(() => {
        dispatch(getUserData())
    }, [])

    return (
        <AppRouter />
    )
}

export default App
