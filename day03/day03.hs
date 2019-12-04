import Data.Char
import Data.List.Split

data Direction = Up | Dwn | Lft | Rgt deriving Eq
data Path = Path Direction Int deriving Eq
data Wire = Wire Int Int deriving Eq

origin :: Wire
origin = Wire 0 0

filename :: String
filename = "input"

getDirection :: String -> Maybe Direction
getDirection s = case s of
    "U" -> Just Up
    "D" -> Just Dwn
    "L" -> Just Lft
    "R" -> Just Rgt
    _   -> Nothing

buildPath :: String -> Int -> Maybe Path
buildPath s i = getDirection s >>= (\x -> pure $ Path x i )

transformInput :: String -> Maybe Path
transformInput s = if length s < 4 && all isDigit sz then Nothing else buildPath dir sz' where
    (dir,sz)  = splitAt 1 s
    sz'       = read sz :: Int

computeWire :: Wire -> [Maybe Path] -> Wire
computeWire  w [] = w
computeWire (Wire x y) (Just (Path Up  size):xs) = computeWire ( Wire  x         (y+size) ) xs
computeWire (Wire x y) (Just (Path Dwn size):xs) = computeWire ( Wire  x         (y-size) ) xs
computeWire (Wire x y) (Just (Path Lft size):xs) = computeWire ( Wire (x-size)    y       ) xs
computeWire (Wire x y) (Just (Path Rgt size):xs) = computeWire ( Wire (x+size)    y       ) xs
computeWire  w         (Nothing:xs)              = computeWire   w                          xs

main :: IO ()
main = do
    content <- readFile filename
    let a = computeWire origin <$> (map transformInput) <$> splitOn "," <$> lines content
    putStrLn ""