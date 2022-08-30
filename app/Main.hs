{-# OPTIONS_GHC -Wno-incomplete-patterns #-}
{-# LANGUAGE LambdaCase #-}
module Main where

import Lib
import Data.Char ( toUpper, toLower, isLetter, isDigit, ord, intToDigit, digitToInt )
import Data.List (union, group, sort, nub, (\\), sortBy, sortOn, elemIndex, delete, find, transpose)
import Data.List.Split (split, startsWithOneOf)
import Data.Ord ( Down(Down) )
import Data.Maybe (fromMaybe, fromJust)

-- Credit Card Mask
maskify :: String -> String
maskify str = replicate len '#' ++ drop len str
    where len = length str - 4
-- maskify str = (\c -> if snd c < length str - 4 then '#' else fst c) <$> zip str [0..]

-- Jaden Casing Strings
toJadenCase :: String -> String
toJadenCase = unwords . map (\(x:xs) -> toUpper x : xs) . words

-- Beginner Series #3 Sum of Numbers
getSum :: Int -> Int -> Int
getSum a b | a > b = sum [b..a]
           | otherwise = sum [a..b]

-- Who likes it?
likes :: [String] -> String
likes [] = "no one likes this"
likes [x] = x ++ " likes this"
likes [x, y] = x ++ " and " ++ y ++ " like this"
likes [x, y, z] = x ++ ", " ++ y ++ " and " ++ z ++ " like this"
likes (x : y : xs) = x ++ ", " ++ y ++ " and " ++ show (length xs) ++ " others like this"

-- Array.diff
difference :: Eq a => [a] -> [a] -> [a]
difference a b = filter (`notElem` b) a

-- N-th Fibonacci
fib :: Int -> Integer
fib n | n >= 1 = go 0 1 n
      | otherwise = toInteger n
  where go a b 1 = a
        go a b n = go b (a + b) (n - 1)

-- Multiples of 3 or 5
solution :: Int -> Int
solution n = sum $ [x | x <- [0 .. n - 1], x `mod` 3 == 0 || x `mod` 5 == 0]
-- solution n = sum $ [3,6..n-1] `union` [5,10..n-1]

-- Bit Counting
countBits :: Int -> Int
countBits n = go 0 n
  where go a 0 = a
        go a n = go (a + n `mod` 2) (n `div` 2)
-- {-# LANGUAGE CPP #-}
-- import Data.Bits {- /* -} hiding {- */ -} (popCount, popCountDefault)
-- countBits :: Int -> Int
-- countBits = popCount

-- Count the number of Duplicates
duplicateCount :: String -> Int
duplicateCount = length . filter ((> 1) . length) . group . sort . map toLower

-- Duplicate Encoder
duplicateEncode :: String -> String
duplicateEncode n = map (\x -> if x `elem` strDub then ')' else '(') str
  where str = map toLower n
        strDub = str \\ nub str

-- Square Every Digit
squareDigit :: Int -> Int
squareDigit x
  | x < 0  = -(squareDigit (-x))
  | x < 10 = x*x
  | otherwise = if r <= 3
                then squareDigit q * 10 + r*r
                else squareDigit q * 100 + r*r
                where (q, r) = x `quotRem` 10

-- Descending Order
descendingOrder :: Integer -> Integer
descendingOrder = read . sortOn Down . show

-- Tribonacci
tribonacci :: Num a => (a, a, a) -> Int -> [a]
tribonacci (a, b, c) 0 = []
tribonacci (a, b, c) n = a : tribonacci (b, c, a + b + c) (n - 1)

-- Take a Ten Minutes Walk
isValidWalk :: [Char] -> Bool
isValidWalk walk = length (take 11 walk) == 10 && count 'w' walk == count 'e' walk && count 'n' walk == count 's' walk
  where count c = length . filter (== c)

-- Unique In Order
uniqueInOrder :: Eq a => [a] -> [a]
uniqueInOrder = map head . group

-- Equal Sides Of An Array
findEvenIndex :: [Int] -> Int
findEvenIndex (a:rr) = go 0 a 0 (sum rr) rr
  where go ix p left right ar | left == right = ix
                              | null ar = -1
                              | (a:as) <- ar = go (ix+1) a (p+left) (right-a) as

-- findEvenIndex = fromMaybe (-1) . elemIndex True .
--   (zipWith (==) <$> scanl1 (+) <*> scanr1 (+))

-- Detect Pangram
isPangram :: String -> Bool
isPangram str | length go == length ['a'..'z'] = True
              | otherwise = False
                where go = nub $ map toLower $ filter isLetter str
-- isPangram str = all (`elem` (map toLower str)) ['a'..'z']

-- Does my number look big in this?
narcissistic :: Integral n => n -> Bool
narcissistic n = n == sum (map (\x -> x^length num) num)
  where num = go [] n
        go a 0 = a
        go a n = go (n `mod` 10 : a) (n `div` 10)

-- Find the unique number
getUnique :: [Int] -> Int
getUnique (x : xs) = case filter (/= x) xs of
  [a] -> a
  _   -> x

-- Title Case
titleCase :: String -> String -> String
titleCase minor title =
  case title of
    "" -> ""
    _  -> unwords $ map (\x -> if x `notElem` a then funUpper x else x) b
  where funUpper (x : xs) = toUpper x : xs
        a = words $ map toLower minor
        b = words $ toUpper (head title) : map toLower (tail title)

-- Sort Numbers
sortNumbers :: [Int] -> Maybe [Int]
sortNumbers [] = Nothing
sortNumbers x = Just $ sort x

-- Moving Zeros To The End
moveZeros :: [Int] -> [Int]
moveZeros (x : xs) | x == 0 = moveZeros xs ++ [x]
                   | otherwise = x : moveZeros xs
-- moveZeros = sortOn (== 0)
-- moveZeros = uncurry (++) . partition (/= 0)
-- moveZeros [] = []

-- Categorize New Member
data Membership = Open | Senior deriving (Eq, Show)
openOrSenior :: [(Int, Int)] -> [Membership]
openOrSenior = map (\(x, y) -> if x >= 55 && y > 7 then Senior else Open)

-- Valid Parentheses
validParentheses :: String -> Bool
validParentheses = go 0
  where go n [] = n == 0
        go n (x:xs) | n < 0 = False
                    | otherwise = if x == '(' then go (n + 1) xs else go (n - 1) xs
-- validParentheses :: String -> Bool
-- validParentheses = null . foldr go []
--   where go '(' (')' : xs) = xs
--         go x xs = x : xs

-- Convert string to camel case
toCamelCase :: String -> String
toCamelCase str = concat $ zipWith (\a b -> if b == 0 then a else capitalize a) (words' str) [0..]
            where capitalize [] = []
                  capitalize (x:xs) = toUpper x : xs
                  words' [] = []
                  words' str = before : words' (dropWhile (\x -> x == '_' || x == '-') after)
                    where (before, after) = break (\x -> x == '_' || x == '-') str
-- toCamelCase = f False
--   where f _ [] = []
--         f _ ('_':xs) = f True xs
--         f _ ('-':xs) = f True xs
--         f False (x:xs) = x : f False xs
--         f True  (x:xs) = toUpper x : f False xs

-- Your order, please
yourOrderPlease :: String -> String
yourOrderPlease = unwords . sortOn (head . filter isDigit) . words

-- Build Tower
buildTower :: Int -> [String]
buildTower n = map (\i -> replicate (n - i - 1) ' ' ++ replicate (2 * i + 1) '*' ++ replicate (n - i - 1) ' ') [0..n - 1]

-- Find the missing letter
findMissingLetter :: [Char] -> Char
findMissingLetter cs = head $ [head cs .. last cs] \\ cs
-- findMissingLetter (x:y:xs) | y == next = findMissingLetter xs
--                            | otherwise = next
--                             where next = succ x

-- Scoring Tests
scoreTest :: (Integral a) => [a] -> a -> a -> a -> a
scoreTest li a b c = sum $
  map (\case
          0 -> a
          1 -> b
          2 -> -c
          _ -> error "WAT!") li

-- Are they the "same"?
comp :: [Integer] -> [Integer] -> Bool
comp as bs = sort (map (^2) as) == sort bs

-- Permutations
permutations :: String -> [String]
permutations "" = [""]
permutations xs = [x : y | x <- nub xs, y <- permutations $ delete x xs]

-- Highest Scoring Word
high :: String -> String
high myStr
  | null myStr = ""
  | otherwise = snd $ fromJust $ find (\(v, _) -> v == fst (last str)) str
  where
    str = sortOn fst $ (\c -> (value c, c)) <$> words myStr
    value = sum . map (\x -> ord x - ord 'a' + 1)

-- Pizza pieces
maxPizza :: Integer -> Maybe Integer
maxPizza n | n < 0 = Nothing
           | otherwise = Just $ n * (n + 1) `div` 2 + 1

-- Highest Rank Number in an Array
highestRank :: Ord c => [c] -> c
highestRank = head . last . sortOn length . group . sort

-- Break camelCase
solution' :: String -> String
solution' = unwords . split (startsWithOneOf ['A'..'Z'])
-- solution` = unwords . groupBy (\x y -> isLower y)

-- Snail
snail :: [[Int]] -> [Int]
snail [] = []
snail (x:xs) = x ++ snail (reverse $ transpose xs)

-- Odd or Even?
oddOrEven :: Num a => Integral a => [a] -> String
oddOrEven xs | odd $ sum xs = "odd"
             | otherwise    = "even"

-- Form The Minimum
minValue :: [Int] -> Int
-- minValue = read . map intToDigit . sort . nub
minValue = foldl1 (\x y -> x * 10 + y) . sort . nub

-- Remove the minimum
removeSmallest :: [Int] -> [Int]
removeSmallest = delete =<< minimum

-- Collatz
collatz :: Int -> String
collatz n
  | n == 1 = "1"
  | otherwise = show n ++ "->" ++ collatz (go n)
  where go n
         | even n = n `div` 2
         | otherwise = n * 3 + 1

-- Is a number prime?
isPrime :: Integer -> Bool
isPrime x
  | x <= 1 = False
  | otherwise = null [i | i <- [2.. floor $ sqrt $ fromInteger x], x `mod` i == 0]

-- Printer Errors
printerError :: [Char] -> [Char]
printerError s = show (length str) ++ "/" ++ show (length s)
    where str = filter (> 'm') s

-- String ends with?
solution'' :: String -> String -> Bool
solution'' s1 s2 = drop len s1 == s2
    where len = length s1 - length s2
-- solution'' = flip isSuffixOf

main :: IO ()
main = someFunc
