import Data.Char (isDigit)
import Control.Monad (sequence)

filename :: String
filename = "input"

countFuel :: Int -> Int
countFuel m = m `div` 3 - 2

countRemainingFuel :: Int -> Int
countRemainingFuel m = countRemainingFuel' m 0 where
    countRemainingFuel' m n = let m' = countFuel m in if m' > 0 then countRemainingFuel' m' n+m' else n

checkInput :: String -> Maybe Int
checkInput s = if all isDigit s then Just (read s::Int) else Nothing

main :: IO ()
main = do
    content <- readFile filename
    let values = fmap sum $ sequence $ fmap countFuel <$> checkInput <$> lines content
    case (values >>= (\x -> pure $ x + countRemainingFuel x)) of
        Just v  -> putStrLn $ show v
        Nothing -> putStrLn "Invalid input"