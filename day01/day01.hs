import Data.Char (isDigit)
import Control.Monad (sequence)

filename :: String
filename = "input"

countFuel :: Int -> Int
countFuel m = let m' = m `div` 3 in m' - 2

checkInput :: String -> Maybe Int
checkInput s = if all isDigit s then Just (read s::Int) else Nothing

main :: IO ()
main = do
    content <- readFile filename
    let values = fmap sum $ sequence $ fmap countFuel <$> checkInput <$> lines content
    case values of
        Just v  -> putStrLn $ show v
        Nothing -> putStrLn "Invalid input"